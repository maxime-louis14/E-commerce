package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	User_id       string             `json:"user_id"`
	Role          string             `json:"role"`
	Nom           string             `json:"nom" validate:"required,min=2,max=100"`
	Prenom        string             `json:"prenom" validate:"required,min=2,max=100"`
	Password      string             `json:"password" validate:"required,min=1"`
	Email         string             `json:"email" validate:"email,required"`
	AvatarURL     *string            `json:"avatarurl"` // Champ pour l'URL de l'avatar
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}
