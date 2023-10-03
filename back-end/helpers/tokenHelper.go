package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/maxime-louis14/Clone-site-Vilebrequin/database"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SignedDetails représente les détails signés du token JWT
type SignedDetails struct {
	Email     string
	Nom       string
	Prenom    string
	Uid       string
	User_type string
	jwt.StandardClaims
}

// Collection d'utilisateurs dans la base de données
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// SECRET_KEY est la clé secrète utilisée pour signer les tokens JWT
var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens génère à la fois le token détaillé et le token de rafraîchissement
func GenerateAllTokens(email string, nom string, prenom string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	// Création des claims du token
	claims := &SignedDetails{
		Email:     email,
		Nom:       nom,
		Prenom:    prenom,
		Uid:       uid,
		User_type: userType,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(), // Token expirera dans 24 heures
		},
	}

	// Création des claims du token de rafraîchissement
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(), // Token de rafraîchissement expirera en 7 jours
		},
	}

	// Création du token en utilisant la méthode de signature HS256 et la clé secrète
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

// ValidateToken valide le token JWT
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	// Analyse du token en utilisant la clé secrète
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}

// UpdateAllTokens renouvelle les tokens de l'utilisateur lors de sa connexion
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	// Mise à jour des tokens et de la date de mise à jour
	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: Updated_at})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	// Mettre à jour le document utilisateur en base de données
	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)
	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}

	return
}
