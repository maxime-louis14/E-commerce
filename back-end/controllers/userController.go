package controllers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/maxime-louis14/Clone-site-Vilebrequin/database"
	"github.com/maxime-louis14/Clone-site-Vilebrequin/helpers"
	models "github.com/maxime-louis14/Clone-site-Vilebrequin/models"
)

var userCollection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

// HashPassword est utilisée pour crypter le mot de passe avant de le stocker dans la base de données
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

// VerifyPassword vérifie le mot de passe saisi en le comparant au mot de passe contenu dans la base de données.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = "login ou mot de passe incorrect"
		check = false
	}
	return check, msg
}

// SignUp est l'API utilisée pour créer un nouvel utilisateur
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Valider les données de l'utilisateur
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Vérifier si l'utilisateur existe déjà
		count, err := userCollection.CountDocuments(ctx, primitive.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erreur lors de la vérification de l'adresse e-mail"})
			return
		}
		defer cancel()

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cette adresse e-mail ou ce numéro de téléphone existe déjà"})
			return
		}

		// Hachage du mot de passe
		password := HashPassword(user.Password)
		user.Password = password

		currentTime := time.Now()
		user.Created_at = currentTime
		user.Updated_at = currentTime

		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		// Génération des tokens
		token, refreshToken, _ := helper.GenerateAllTokens(user.Email, user.Nom, user.Prenom, user.User_id, "")
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := "L'utilisateur n'a pas pu être créé"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

// Login est l'API utilisée pour authentifier un utilisateur
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, primitive.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login ou mot de passe incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(user.Password, foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(foundUser.Email, foundUser.Nom, foundUser.Prenom, foundUser.User_id, "")

		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		c.JSON(http.StatusOK, foundUser)
	}
}

func UploadAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, uidExists := c.Get("uid")

		// Assurez-vous que l'UID existe avant de l'utiliser
		if !uidExists {
			log.Println("UID de l'utilisateur manquant")
			c.JSON(http.StatusBadRequest, gin.H{"error": "UID de l'utilisateur manquant"})
			return
		}

		// Récupérer le fichier d'avatar à partir de la requête HTTP
		file, err := c.FormFile("avatar")
		if err != nil {
			log.Println("Erreur lors de la récupération du fichier d'avatar :", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erreur lors de la récupération du fichier d'avatar"})
			return
		}

		// Vérifier l'extension du fichier
		ext := filepath.Ext(file.Filename)
		if !isValidExtension(ext) {
			log.Println("Extension de fichier non autorisée :", ext)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Extension de fichier non autorisée"})
			return
		}

		// Générer un nom de fichier unique pour l'image d'avatar
		avatarFileName := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)

		// Définir le chemin complet du fichier sur le serveur
		fullPath := "public/uploads/avatars/" + avatarFileName

		// Lire le contenu du fichier en tant que tableau de bytes
		fileContent, err := file.Open()
		if err != nil {
			log.Println("Erreur lors de l'ouverture du fichier source :", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'ouverture du fichier source"})
			return
		}
		defer fileContent.Close()

		data, err := io.ReadAll(fileContent)
		if err != nil {
			log.Println("Erreur lors de la lecture du contenu du fichier :", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la lecture du contenu du fichier"})
			return
		}

		// Sauvegarder l'image sur le serveur
		if err := SaveImageToFile(fullPath, data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la sauvegarde de l'image sur le serveur"})
			return
		}

		// Définir l'URL de l'avatar
		avatarURL := "public/uploads/avatars/" + avatarFileName

		// Créer une instance de la structure User
		avatar := &models.User{
			User_id:           uid.(string),
			AvatarImageData:   data,
			AvatarContentType: file.Header.Get("Content-Type"),
			AvatarURL:         avatarURL,
		}

		// Insérer l'avatar dans la base de données
		err = UpdateUserAvatar(uid.(string), avatar)
		if err != nil {
			log.Println("Erreur lors de la mise à jour de l'avatar dans la base de données :", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de l'avatar dans la base de données"})
			return
		}

		log.Println("Avatar mis à jour avec succès")

		// Répondre avec un succès
		c.JSON(http.StatusOK, gin.H{"message": "Avatar téléversé avec succès"})
	}
}

func SaveImageToFile(filePath string, data []byte) error {
	// Assurez-vous que le répertoire de destination existe
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		log.Printf("Erreur lors de la création du répertoire : %v\n", err)
		return err
	}

	// Écrire le contenu de l'image dans le fichier
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Printf("Erreur lors de l'écriture du fichier : %v\n", err)
		return err
	}

	log.Printf("Image enregistrée avec succès : %s\n", filePath)
	return nil
}

// isValidExtension vérifie si l'extension du fichier est autorisée
func isValidExtension(ext string) bool {
	allowedExtensions := map[string]bool{
		".gif":  true,
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	return allowedExtensions[ext]
}

// Mettre à jour les champs d'avatar en utilisant UpdateOne
func UpdateUserAvatar(user_id string, user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filtre pour trouver l'utilisateur par ID (utilisez "user_id")
	filter := bson.M{"user_id": user_id}

	// Contrôle des champs d'avatar
	if user.AvatarImageData == nil || user.AvatarContentType == "" || user.AvatarURL == "" {
		// Vous pouvez choisir de gérer cette situation comme vous le souhaitez, par exemple, retourner une erreur ou effectuer une action par défaut.
		return errors.New("erreur lors de la récupération du fichier d'avatar")
	}

	// Définition de la mise à jour pour mettre à jour les champs d'avatar
	update := bson.M{
		"$set": bson.M{
			"avatarimagedata":   user.AvatarImageData,
			"avatarcontenttype": user.AvatarContentType,
			"avatarurl":         user.AvatarURL,
		},
	}

	// Effectuer la mise à jour dans la base de données en utilisant UpdateOne
	_, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		// Ajoutez un message de journalisation en cas d'erreur
		log.Println("Erreur lors de la mise à jour de l'avatar dans la base de données :", err)
		return err
	}

	return nil
}

// Mettre à jour les champs Avatar et ContentType en utilisant UpdateOne
func UpdateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filtre pour trouver l'utilisateur par ID (utilisez "user_id")
	filter := bson.M{"user_id": user.User_id}

	// Contrôle des champs d'avatar
	if user.AvatarURL == "" || user.AvatarContentType == "" {
		// Vous pouvez choisir de gérer cette situation comme vous le souhaitez, par exemple, retourner une erreur ou effectuer une action par défaut.
		return errors.New("les champs d'avatar ne peuvent pas être vides")
	}

	// Définition de la mise à jour pour mettre à jour les champs Avatar et ContentType
	update := bson.M{
		"$set": bson.M{
			"avatarurl":         user.AvatarURL,
			"avatarcontenttype": user.AvatarContentType,
		},
	}

	// Effectuer la mise à jour dans la base de données en utilisant UpdateOne
	_, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		// Ajoutez un message de journalisation en cas d'erreur
		log.Println("Erreur lors de la mise à jour de l'utilisateur dans la base de données:", err)
		return err
	}

	return nil
}

func GetAvatarImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Récupérer l'ID de l'utilisateur depuis le contexte
		userID, userIDExists := c.Get("uid")

		if !userIDExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'utilisateur manquant"})
			return
		}

		// Rechercher l'utilisateur par son ID
		user, err := FindUserByID(userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la recherche de l'utilisateur"})
			return
		}

		if user.AvatarURL == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Aucun avatar trouvé pour cet utilisateur"})
			return
		}

		// Construire le chemin complet du fichier d'avatar
		avatarFilePath := filepath.Join("public/uploads/avatars", filepath.Base(user.AvatarURL))

		// Lire le contenu de l'image depuis la base de données ou le système de fichiers
		imageData, err := ioutil.ReadFile(avatarFilePath)
		if err != nil {
			log.Println("Erreur lors de la lecture du fichier d'avatar :", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la lecture du fichier d'avatar"})
			return
		}

		// Utiliser le type de contenu préféré par le client s'il est spécifié, sinon utiliser le type de contenu de l'utilisateur
		contentType := c.ContentType()
		if contentType == "" {
			contentType = user.AvatarContentType
		}

		// Envoyer le contenu de l'image d'avatar au front-end avec le bon type de contenu
		c.Data(http.StatusOK, contentType, imageData)
	}
}

// FindUserByID recherche un utilisateur dans la base de données par son ID
func FindUserByID(user_id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	filter := bson.M{"user_id": user_id}

	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
