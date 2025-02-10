package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nntruong112/docker-practice/controllers"
)

func CarRoutes(router *gin.Engine) {
	router.POST("/cars", controllers.CreateCar)
	router.GET("/cars", controllers.GetCars)
	router.GET("/cars/:id", controllers.GetCarByID)
	router.PUT("/cars/:id", controllers.UpdateCar)
	router.DELETE("/cars/:id", controllers.DeleteCar)
}
