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
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
}
