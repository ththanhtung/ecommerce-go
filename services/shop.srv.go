package services

import (
	"ecom/helpers"
	"ecom/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

// roles
const (
	SHOP   string = "001"
	WRITER string = "002"
	EDITOR string = "003"
	ADMIN  string = "004"
)

type shopRepo interface {
	CreateNewShop(*models.ShopCreateRequest) (*models.Shop, error)
	CountShop(filter any) (int, error)
}

type ShopService struct {
	shopRepo     shopRepo
}

func NewShopService(shopRepo shopRepo) *ShopService {
	return &ShopService{
		shopRepo:     shopRepo,
	}
}

func (s *ShopService) Signup(shop *models.ShopCreateRequest) (*models.Shop, error) {
	shopCount, err := s.shopRepo.CountShop(bson.D{{"email", shop.Email}})
	if err != nil {
		return nil, err
	}
	if shopCount > 0 {
		return nil, errors.New("user already existed")
	}

	// hashing password
	hashedPassword, err := helpers.HashPassword(shop.Password)
	if err != nil {
		return nil, err
	}

	shop.Password = hashedPassword

	newShop, err := s.shopRepo.CreateNewShop(shop)
	if err != nil {
		return nil, err
	}

	return newShop, nil
}
