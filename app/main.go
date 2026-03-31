package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb://mongo:27017"

type User struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	collection := client.Database("trivy_demo").Collection("users")

	router := gin.Default()
	router.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertCtx, insertCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer insertCancel()

		result, err := collection.InsertOne(insertCtx, bson.M{
			"username": user.Username,
			"email":    user.Email,
			"created":  time.Now().UTC(),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
	})

	router.GET("/users", func(c *gin.Context) {
		findCtx, findCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer findCancel()

		cursor, err := collection.Find(findCtx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(findCtx)

		var users []User
		if err := cursor.All(findCtx, &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
