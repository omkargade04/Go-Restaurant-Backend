package controllers

import (
	"context"
	"fmt"
	"go-restro-backend/database"
	"go-restro-backend/models"
	"math"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error occured while retreiving food"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food

		if err := c.BindJSON(&food); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		
		if validationErr := validate.Struct(food); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": validationErr.Error()})
			return
		}

		if err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu); err != nil {
			msg := fmt.Sprintf("Menu not found")
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
			return
		}

		food.Created_at = time.Now()
		food.Updated_at = time.Now()
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()

		if food.Price != nil {
			num := toFixed(*food.Price, 2)
			food.Price = &num
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Price is required"})
			return
		}

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			msg := fmt.Sprintf("Food item not created")
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func round(num float64) int {
	return int(num + 0.5)
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
