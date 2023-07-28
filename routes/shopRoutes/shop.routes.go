package shoproutes

import "github.com/gin-gonic/gin"

type shopController interface {
	ShopSignUp() gin.HandlerFunc
}

func RegistShopRoutes(r *gin.RouterGroup, controller shopController){
	shopRoutes := r.Group("/shop")
	PublicShopRoutes(shopRoutes, controller)
}

func PublicShopRoutes(r *gin.RouterGroup, controller shopController){
	r.POST("/auth/signup", controller.ShopSignUp())
}

func PrivateShopRoutes(){

}