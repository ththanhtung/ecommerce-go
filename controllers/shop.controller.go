package controllers

import (
	"ecom/helpers"
	"ecom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type shopService interface {
	Signup(*models.ShopCreateRequest) (*models.Shop, error)
}

type KeyTokenService interface {
	CreateNewKeyToken(keytoken *models.KeyTokenCreateRequest) (*models.KeyToken, error)
}

type shopController struct {
	shopSrv     shopService
	keyTokenSrv KeyTokenService
}

func NewShopController(shopSrv shopService, keytokenSrv KeyTokenService) *shopController {
	return &shopController{
		shopSrv:     shopSrv,
		keyTokenSrv: keytokenSrv,
	}
}

func (s *shopController) ShopSignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newShopRequest *models.ShopCreateRequest
		if err := c.ShouldBindJSON(&newShopRequest); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error:": err.Error(),
			})
			return
		}

		newShop, err := s.shopSrv.Signup(newShopRequest)

		privateKeyPEM, publicKeyPEM, err := helpers.GenerateKeyPair()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error:": err.Error(),
			})
			return
		}

		accessToken, refreshToken, err := helpers.CreateTokensPair(&helpers.JwtClaims{
			UserID: newShop.ID.Hex(),
			Email:  newShop.Email,
		}, privateKeyPEM, publicKeyPEM)

		keyToken := &models.KeyTokenCreateRequest{
			UserID:      newShop.ID,
			PrivateKey:  string(privateKeyPEM),
			PublicKey:   string(publicKeyPEM),
			RefeshToken: refreshToken,
		}

		s.keyTokenSrv.CreateNewKeyToken(keyToken)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		tokens := models.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		shopResponse := models.ShopResponse{
			ID:    newShop.ID.Hex(),
			Email: newShop.Email,
			Name:  newShop.Name,
			Roles: newShop.Roles,
		}

		response := &models.ShopCreatedResponse{
			Shop:   shopResponse,
			Tokens: tokens,
		}

		c.JSON(http.StatusCreated, response)
	}
}
