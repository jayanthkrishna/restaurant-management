package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jayanthkrishna/restaurant-management-golang/database"
	"github.com/jayanthkrishna/restaurant-management-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		orderId := c.Param("order_id")

		var order models.Order

		err := foodCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the order item"})
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(order)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()

			if err != nil {
				msg := fmt.Sprintf("message: Table was not found")

				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

			order.ID = primitive.NewObjectID()
			order.Order_id = order.ID.Hex()

			res, insertErr := orderCollection.InsertOne(ctx, order)

			if insertErr != nil {
				msg := fmt.Sprintf("Order Item not created")

				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}

			defer cancel()

			c.JSON(http.StatusOK, res)
		}
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var table models.Table
		var order models.Order

		var updatedObj primitive.D

		orderId := c.Param("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()

			if err != nil {
				msg := fmt.Sprintf("message: Table was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			updatedObj = append(updatedObj, bson.E{"table", order.Table_id})

		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		updatedObj = append(updatedObj, bson.E{"updated_at", order.Updated_at})

		upsert := true

		filter := bson.M{"order_id": orderId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		res, err := orderCollection.UpdateOne(ctx, filter, bson.D{
			{"$set", updatedObj},
		},
			&opt)

		if err != nil {
			msg := fmt.Sprintf("Order Item update failed")

			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, res)
	}
}

func OrderItemOrderCreator(order models.Order) string {
	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	orderCollection.InsertOne(context.TODO(), order)

	return order.Order_id
}
