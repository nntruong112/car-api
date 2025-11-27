package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/nntruong112/docker-practice/database"
	_ "github.com/nntruong112/docker-practice/docs" // Swagger docs
	"github.com/nntruong112/docker-practice/routes"
	swaggerFiles "github.com/swaggo/files" // <-- correct package
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Car API
// @version 1.0
// @description API for managing cars
// @host localhost:8080
// @BasePath /
func main() {
	client := database.ConnectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()
	routes.CarRoutes(router)

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
