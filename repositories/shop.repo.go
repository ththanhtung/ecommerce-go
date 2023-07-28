package repositories

import (
	"context"
	"ecom/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShopRepo struct {
	collection *mongo.Collection
}

func NewShopRepo(mongoCollection *mongo.Collection) *ShopRepo {
	return &ShopRepo{
		collection: mongoCollection,
	}
}

func (r *ShopRepo) CreateNewShop(newShop *models.ShopCreateRequest) (*models.Shop, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	shop := &models.Shop{
		ID:       primitive.NewObjectID(),
		Email:    newShop.Email,
		Password: newShop.Password,
		Name:     newShop.Name,
		Status:   "active",
		Verify:   false,
		Roles:    []string{"001"},
	}

	createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	shop.CreatedAt = createdAt
	shop.UpdatedAt = updatedAt

	_, err := r.collection.InsertOne(ctx, &shop)
	defer cancel()

	if err != nil {
		return nil, err
	}

	return shop, nil
}

func (r *ShopRepo) CountShop(filter any) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	numberOfShop, err := r.collection.CountDocuments(ctx, filter)
	defer cancel()

	if (err != nil) {
		return 0, err
	}

	return int(numberOfShop), nil
}