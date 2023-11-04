package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/maxime-louis14/Clone-site-Vilebrequin/database"
	helper "github.com/maxime-louis14/Clone-site-Vilebrequin/helpers"
	"github.com/maxime-louis14/Clone-site-Vilebrequin/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
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
		msg = fmt.Sprintf("login ou mot de passe incorrect")
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
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
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
			msg := fmt.Sprintf("L'utilisateur n'a pas pu être créé")
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

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
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

// UploadAvatar gère les téléversements d'avatar
func UploadAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, uidExists := c.Get("uid")

		// Vérifiez si l'UID existe avant de l'utiliser
		// Ajoute un journal pour vérifier l'UID extrait
		log.Printf("UID de l'utilisateur lors de la génération des tokens : %v", uid)

		// Assurez-vous que l'UID existe avant de l'utiliser
		if !uidExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "UID de l'utilisateur manquant"})
			return
		}

		// Récupérer le fichier d'avatar à partir de la requête HTTP
		file, err := c.FormFile("avatar")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erreur lors de la récupération du fichier d'avatar"})
			return
		}

		// Vérifier l'extension du fichier (autoriser uniquement .gif, .jpg, .png)
		allowedExtensions := map[string]bool{
			".gif": true,
			".jpg": true,
			".png": true,
		}
		ext := filepath.Ext(file.Filename)
		if !allowedExtensions[ext] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Extension de fichier non autorisée"})
			return
		}

		// Générer un nom de fichier unique pour l'image d'avatar
		avatarFileName := fmt.Sprintf("public/uploads/avatars/%d-%s", time.Now().Unix(), file.Filename)

		// Définir l'URL de l'avatar
		avatarURL := "/public/uploads/avatars/" + avatarFileName

		// Créer le fichier de destination sur le serveur
		outFile, err := os.Create(avatarFileName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création du fichier d'avatar"})
			return
		}
		defer outFile.Close()

		// Copier le contenu du fichier source dans le fichier de destination
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'ouverture du fichier source"})
			return
		}
		defer src.Close()

		_, err = io.Copy(outFile, src)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la copie du contenu du fichier"})
			return
		}

		// Mettre à jour le champ AvatarURL dans la base de données
		user := &models.User{
			User_id:   uid.(string),
			AvatarURL: &avatarURL,
		}

		// Vérifiez si l'UID existe avant de l'utiliser
		if !uidExists {
			log.Println("UID de l'utilisateur manquant")
			c.JSON(http.StatusBadRequest, gin.H{"error": "UID de l'utilisateur manquant"})
			return
		}

		if err := UpdateUser(user); err != nil {
			// Ajoutez un journal pour afficher la valeur de l'UID avant la mise à jour
			log.Println("UID de l'utilisateur avant la mise à jour :", user.User_id)
			log.Println("Erreur lors de la mise à jour de l'avatar dans la base de données:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de l'avatar dans la base de données"})
			return
		}

		if err := UpdateUser(user); err != nil {
			log.Println("Erreur lors de la mise à jour de l'avatar dans la base de données:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de l'avatar dans la base de données"})
			return
		}

		log.Println("Avatar URL mis à jour avec succès :", *user.AvatarURL)

		// Répondre avec un succès et l'URL de l'avatar
		c.JSON(http.StatusOK, gin.H{"message": "Avatar téléversé avec succès", "avatarurl": avatarURL})
	}

}

// UpdateUser met à jour un utilisateur dans la base de données
func UpdateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filtre pour trouver l'utilisateur par ID (utilisez "user_id")
	filter := bson.M{"user_id": user.User_id}

	// Définition de la mise à jour pour mettre à jour le chemin de l'avatar
	update := bson.M{
		"$set": bson.M{
			"avatarurl": user.AvatarURL,
		},
	}

	// Avant de lancer la requête FindOneAndUpdate, ajoutez ces journaux pour vérifier le filtre
	log.Println("Début de la mise à jour de l'utilisateur dans la base de données")
	log.Println("Filtre de recherche :", filter)
	log.Println("Données de l'utilisateur à mettre à jour :", user.User_id)

	// Effectuer la mise à jour dans la base de données en utilisant FindOneAndUpdate
	result := userCollection.FindOneAndUpdate(ctx, filter, update)

	if result.Err() != nil {
		// Ajoutez un message de journalisation en cas d'erreur
		log.Println("Erreur lors de la mise à jour de l'utilisateur dans la base de données:", result.Err())
		return result.Err()
	}

	// Ajouter un message de journalisation pour indiquer la fin de la mise à jour
	log.Println("Fin de la mise à jour de l'utilisateur dans la base de données")

	return nil
}
