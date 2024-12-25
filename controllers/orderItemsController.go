package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc{
	return func(c *gin.Context) {

	}
}

func GetOrderItem() gin.HandlerFunc{
	return func(c *gin.Context) {

	}
}

func GetOrderItemsByOrder() gin.HandlerFunc{
	return func(c *gin.Context) {

	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	
}

func CreateOrderItems() gin.HandlerFunc{
	return func(c *gin.Context) {

	}
}

func UpdateOrderItems() gin.HandlerFunc{
	return func(c *gin.Context) {

	}
}