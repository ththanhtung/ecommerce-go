package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyToken struct {
	ID              primitive.ObjectID `bson:"_id"`
	UserID          primitive.ObjectID `bson:"userId"`
	PublicKey       string             `bson:"publicKey"`
	PrivateKey      string             `bson:"privateKey"`
	RefeshToken     string             `bson:"refeshToken"`
	RefeshTokenUsed []string           `bson:"refeshTokensUsed"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
}

type KeyTokenCreateRequest struct {
	UserID          primitive.ObjectID `bson:"userId"`
	PublicKey       string             `bson:"publicKey"`
	PrivateKey      string             `bson:"privateKey"`
	RefeshToken     string             `bson:"refeshToken"`
}