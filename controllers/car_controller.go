package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nntruong112/docker-practice/database"
	"github.com/nntruong112/docker-practice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateCar godoc
// @Summary Create a new car
// @Description Create a new car entry
// @Tags cars
// @Accept json
// @Produce json
// @Param car body models.Car true "Car info"
// @Success 201 {object} models.Car
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cars [post]
func CreateCar(c *gin.Context) {
	var car models.Car
	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	car.ID = primitive.NewObjectID()
	_, err := database.CarCollection.InsertOne(context.TODO(), car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}
	c.JSON(http.StatusCreated, car)
}

// GetCars godoc
// @Summary List all cars
// @Description Get all cars
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {array} models.Car
// @Failure 500 {object} map[string]string
// @Router /cars [get]
func GetCars(c *gin.Context) {
	var cars []models.Car
	cursor, err := database.CarCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cars"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var car models.Car
		if err := cursor.Decode(&car); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding car"})
			return
		}
		cars = append(cars, car)
	}
	c.JSON(http.StatusOK, cars)
}

// GetCarByID godoc
// @Summary Get a car by ID
// @Description Get a single car by its ID
// @Tags cars
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {object} models.Car
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /cars/{id} [get]
func GetCarByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var car models.Car

	err = database.CarCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&car)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}
	c.JSON(http.StatusOK, car)
}

// UpdateCar godoc
// @Summary Update a car
// @Description Update a car by ID
// @Tags cars
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Param car body models.Car true "Car info"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cars/{id} [put]
func UpdateCar(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var updatedCar models.Car
	if err := c.BindJSON(&updatedCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update := bson.M{"$set": updatedCar}
	_, err = database.CarCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car updated successfully"})
}

// DeleteCar godoc
// @Summary Delete a car
// @Description Delete a car by ID
// @Tags cars
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cars/{id} [delete]
func DeleteCar(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	_, err = database.CarCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete car"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
