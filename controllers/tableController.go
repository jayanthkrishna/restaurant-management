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

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		res, err := tableCollection.Find(ctx, bson.M{})

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing table items"})
		}
		var allTables []bson.M

		if err = res.All(ctx, &allTables); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allTables)

	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		tableId := c.Param("table_id")

		var table models.Table

		err := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the order item"})
		}
		c.JSON(http.StatusOK, table)
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var table models.Table

		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(table)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		table.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		table.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		table.ID = primitive.NewObjectID()
		table.Table_id = table.ID.Hex()

		res, insertErr := tableCollection.InsertOne(ctx, table)

		if insertErr != nil {
			msg := fmt.Sprintf("Table Item not created")

			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, res)

	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var table models.Table

		tableId := c.Param("table_id")

		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}

		var updatedObj primitive.D

		if table.Number_of_guests != nil {
			updatedObj = append(updatedObj, bson.E{"number_of_guests", table.Number_of_guests})
		}

		if table.Table_number != nil {
			updatedObj = append(updatedObj, bson.E{"table_number", table.Table_number})
		}
		table.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		updatedObj = append(updatedObj, bson.E{"updated_at", table.Updated_at})

		upsert := true

		filter := bson.M{"order_id": tableId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		res, err := tableCollection.UpdateOne(ctx, filter, bson.D{
			{"$set", updatedObj},
		},
			&opt)

		if err != nil {
			msg := fmt.Sprintf("Table update failed")

			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, res)
	}
}
