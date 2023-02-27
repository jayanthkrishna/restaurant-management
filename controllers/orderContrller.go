package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jayanthkrishna/restaurant-management-golang/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		res, err := orderCollection.Find(ctx, bson.M{})

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while ordering items"})
		}
		var allOrders []bson.M

		if err = res.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)

	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
