package main

import (
	"ecom/configs"
	"ecom/controllers"
	"ecom/database"
	"ecom/repositories"
	shoproutes "ecom/routes/shopRoutes"
	"ecom/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConfig, serverConfig := configs.LoadDBConfig(".env")
	log.Println(dbConfig)
	// open database
	db := database.NewMongoDB(dbConfig)

	// open collections
	shopCollection := db.OpenCollection("shops")

	// init repos
	shopRepo := repositories.NewShopRepo(shopCollection)

	// init services
	shopService := services.NewShopService(shopRepo)


	shopController := controllers.NewShopController(shopService)

	// init server
	router := gin.Default()

	// api v1
	apiV1Router := router.Group("/v1/api")

	// init routes
	shoproutes.RegistShopRoutes(apiV1Router, shopController)

	router.Run(":"+serverConfig.Port)
}