package controllers

import (
	"ecom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type shopService interface {
	Signup(*models.ShopCreateRequest) (*models.Shop, error)
}

type shopController struct {
	srv shopService
}

func NewShopController(srv shopService) *shopController {
	return &shopController{
		srv: srv,
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

		newShop, err := s.srv.Signup(newShopRequest)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, newShop)
	}
}
