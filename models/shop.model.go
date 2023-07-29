package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shop struct {
	ID        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Name      string             `bson:"name"`
	Status    string             `bson:"status"`
	Verify    bool               `bson:"verify"`
	Roles     []string           `bson:"roles"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type ShopCreateRequest struct {
	Email    string
	Password string
	Name     string
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ShopResponse struct {
	ID    string   `json:"userId"`
	Email string   `json:"email"`
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

type ShopCreatedResponse struct {
	Shop   ShopResponse `json:"shop"`
	Tokens Tokens       `json:"tokens"`
}
