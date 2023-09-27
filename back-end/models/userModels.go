package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID            primitive.ObjectID `bson:"_id"`
    Nom           *string            `json:"nom" validate:"required,min=2,max=100"`
    Prenom        *string            `json:"prenom" validate:"required,min=2,max=100"`
    Password      *string            `json:"password" validate:"required,min=1"`
    Email         *string            `json:"email" validate:"email,required"`
    Token         *string            `json:"token"`
    Refresh_token *string            `json:"refresh_token"`
    Created_at    time.Time          `json:"created_at"`
    Updated_at    time.Time          `json:"updated_at"`
    User_id       string             `json:"user_id"`
    Image         *string            `json:"image"` // Champ pour le chemin/nom de l'avatar
}
