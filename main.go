package main

import (
	"golang-Restaurant/database"
	"golang-Restaurant/middleware"
	"golang-Restaurant/routes"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderItemRoutes(router)
	routes.TableRoutes(router)
	router.Run(":" + port)

}
