package repositories

import (
	"context"
	"ecom/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type KeyTokenRepo struct {
	collection *mongo.Collection
}

func NewKeyTokenRepo(collection *mongo.Collection)*KeyTokenRepo {
	return &KeyTokenRepo{
		collection: collection,
	}
}

func (r *KeyTokenRepo) CreateNewKeyToken(keytoken *models.KeyTokenCreateRequest) (*models.KeyToken, error) {
	newKeytoken := &models.KeyToken{
		ID: primitive.NewObjectID(),
		UserID: keytoken.UserID,
		PublicKey: keytoken.PublicKey,
		PrivateKey: keytoken.PrivateKey,
		RefeshToken: keytoken.RefeshToken,
		RefeshTokenUsed: []string{},
	}
	
	createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	newKeytoken.CreatedAt = createdAt
	newKeytoken.UpdatedAt = updatedAt

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := r.collection.InsertOne(ctx, &newKeytoken)
	defer cancel()
	if err != nil {
		return &models.KeyToken{}, err
	}

	return newKeytoken, nil
}