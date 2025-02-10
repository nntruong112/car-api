package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/nntruong112/docker-practice/database"
	"github.com/nntruong112/docker-practice/routes"
)

func main() {
	client := database.ConnectDB()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()
	routes.CarRoutes(router)

	router.Run(":8080")
}
