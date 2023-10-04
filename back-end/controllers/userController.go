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

// HashPassword is used to encrypt the password before it is stored in the DB
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
		msg = fmt.Sprintf("login or password is incorrect")
		check = false
	}
	return check, msg
}

// CreateUser is the api used to get a single user
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		fmt.Println(validationErr)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the email"})
			return
		}
		defer cancel()

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exists"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		currentTime := time.Now()
		user.Created_at = currentTime
		user.Updated_at = currentTime

		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.Nom, *user.Prenom, user.User_id, "")
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

// Login is the api used to get a single user
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.Nom, *foundUser.Prenom, foundUser.User_id, "")

		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		c.JSON(http.StatusOK, foundUser)
	}
}

// UploadAvatar gère les téléversements d'avatar
func UploadAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtenir le token JWT à partir de l'en-tête de la demande
		tokenClient := c.GetHeader("Authorization")

		// Rechercher l'utilisateur dans la base de données en utilisant le token client
		user, err := FindUserByToken(tokenClient)

		if err != nil {
			// En cas d'erreur lors de la recherche de l'utilisateur, renvoyer une réponse 500
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la recherche de l'utilisateur", "message": "Erreur interne du serveur"})
			return
		}

		// Si l'utilisateur n'est pas trouvé, renvoyer une réponse 404
		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur introuvable", "message": "404"})
			return
		}

		// Vérifier si le token côté client correspond au token stocké en base de données
		if user.Token != nil && *user.Token != tokenClient {
			// En cas de token client incorrect, renvoyer une réponse 401
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token client incorrect", "message": "Non autorisé"})
			return
		}

		// Récupérer le fichier d'avatar à partir de la requête HTTP
		file, err := c.FormFile("avatar")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erreur lors de la récupération du fichier d'avatar"})
			return
		}

		// Vérifier l'extension du fichier (autoriser uniquement .gif et .jpg)
		allowedExtensions := map[string]bool{
			".gif": true,
			".jpg": true,
		}
		ext := filepath.Ext(file.Filename)
		if !allowedExtensions[ext] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Extension de fichier non autorisée"})
			return
		}

		// Générer un nom de fichier unique pour l'image d'avatar
		avatarFileName := fmt.Sprintf("public/uploads/%d-%s", time.Now().Unix(), file.Filename)

		// Définir l'URL de l'avatar
		avatarURL := "/public/uploads/avatars" + avatarFileName

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
		user.AvatarURL = &avatarURL

		if err := UpdateUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de l'avatar dans la base de données"})
			return
		}

		// Répondre avec un succès et l'URL de l'avatar
		c.JSON(http.StatusOK, gin.H{"message": "Avatar téléversé avec succès", "avatar_url": avatarURL})
	}
}

// FindUserByToken recherche un utilisateur par token dans la base de données
func FindUserByToken(token string) (*models.User, error) {
	var user models.User

	// Ajoutez un message de log pour indiquer le début de la fonction
	log.Println("Début de la recherche de l'utilisateur par token")

	// Utiliser la connexion MongoDB pour rechercher l'utilisateur par token
	err := userCollection.FindOne(context.TODO(), bson.M{"token": token}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Ajoutez un message de log en cas de document non trouvé
			log.Println("Aucun document trouvé pour le token :", token)
			return nil, nil // Aucun document trouvé, retourne nil sans erreur
		}
		// En cas d'autres erreurs, ajoutez un message de log d'erreur
		log.Println("Erreur lors de la recherche de l'utilisateur :", err)
		return nil, err
	}
	// Ajoutez un message de log pour indiquer la fin de la fonction
	log.Println("Fin de la recherche de l'utilisateur par token")

	return &user, nil
}

// UpdateUser met à jour un utilisateur dans la base de données
func UpdateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Filtre pour trouver l'utilisateur par ID
	filter := bson.M{"user_id": user.User_id}
	// Définition de la mise à jour pour mettre à jour le chemin de l'avatar
	update := bson.M{
		"$set": bson.M{
			"avatar": user.Avatar,
		},
	}
	// Effectuer la mise à jour dans la base de données
	_, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
