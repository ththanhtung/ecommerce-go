package services

import "ecom/models"

type KeyTokenRepo interface {
	CreateNewKeyToken(keytoken *models.KeyTokenCreateRequest) (*models.KeyToken, error)
}

type KeyTokenService struct {
	keyTokenRepo KeyTokenRepo
}

func NewKeyTokenService(keyTokenRepo KeyTokenRepo) *KeyTokenService {
	return &KeyTokenService{
		keyTokenRepo: keyTokenRepo,
	}
}

func (s *KeyTokenService) CreateNewKeyToken(keytoken *models.KeyTokenCreateRequest) (*models.KeyToken, error){
	return s.keyTokenRepo.CreateNewKeyToken(keytoken)
}