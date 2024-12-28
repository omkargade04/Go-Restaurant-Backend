package controllers

import (
	"context"
	"fmt"
	"go-restro-backend/database"
	"go-restro-backend/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error while retreiving order items"})
			return
		}
		var allOrder []bson.M
		err = result.All(ctx, &allOrder)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrder)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("order_id")
		var order models.Order

		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error occured while retreiving order"})
		}
		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table models.Table
		var order models.Order

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		validateErr := validate.Struct(order)

		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": validateErr.Error()})
			return
		}

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message:Table not found")
				c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
				return
			}

			order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			order.ID = primitive.NewObjectID()
			order.Order_id = order.ID.Hex()

			result, insertErr := orderCollection.InsertOne(ctx, order)

			if insertErr != nil {
				msg := fmt.Sprintf("Food item not created")
				c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
				return
			}
			defer cancel()
			c.JSON(http.StatusOK, result)
		}
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table models.Table
		var order models.Order

		var updateObj primitive.D

		orderId := c.Param("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message: Table was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
				return
			}
			updateObj = append(updateObj, bson.E{Key: "table", Value: order.Table_id})

			order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{Key: "updated_at", Value: order.Updated_at})

			upsert := true

			filter := bson.M{"order_id": orderId}
			opt := options.UpdateOptions{
				Upsert: &upsert,
			}

			result, err := orderCollection.UpdateOne(
				ctx,
				filter,
				bson.D{
					{Key: "$set", Value: updateObj},
				},
				&opt,
			)

			if err != nil {
				msg := fmt.Sprintf("Order item update failed")
				c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
				return
			}
			defer cancel()
			c.JSON(http.StatusOK, result)

		}
	}
}
