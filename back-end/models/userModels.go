package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id"`
	User_id           string             `json:"user_id"`
	Nom               string             `json:"nom" validate:"required,min=2,max=100"`
	Prenom            string             `json:"prenom" validate:"required,min=2,max=100"`
	Password          string             `json:"password" validate:"required,min=1"`
	Email             string             `json:"email" validate:"email,required"`
	AvatarURL         string             `json:"avatarurl"`
	AvatarImageData   []byte             `json:"avatarimagedata"`   // Ajoutez le champ pour les données de l'avatar
	AvatarContentType string             `json:"avatarcontenttype"` // Ajoutez le champ pour le type de contenu de l'avatar
	Role              string             `json:"role"`
	Token             *string            `json:"token"`
	Refresh_token     *string            `json:"refresh_token"`
	Created_at        time.Time          `json:"created_at"`
	Updated_at        time.Time          `json:"updated_at"`
}
