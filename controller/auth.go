package controller

import (
	"context"
	"net/http"
	"testapi/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client     *mongo.Client     = db.Client
	collection *mongo.Collection = client.Database("testapi").Collection("user_table")
	cache      *redis.Client     = db.RDB
	ctx        context.Context   = context.Background()
)

//SignUp ...
func SignUp(c *gin.Context) {
	username := c.Query("username")
	user := db.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: c.Query("password"),
	}
	var e db.User
	err := collection.FindOne(context.TODO(), bson.D{}).Decode(&e)
	if err == mongo.ErrNoDocuments {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "user created"})
		return
	}
	c.JSON(http.StatusAlreadyReported, gin.H{"message": "User already exists!"})
	return
}

//SignIn ...
func SignIn(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user db.User
	filter := bson.D{
		primitive.E{Key: "username", Value: username},
	}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusOK, gin.H{"message": "Database error"})
	}
	if user.Password != password {
		c.JSON(http.StatusOK, gin.H{"message": "Invalid Credetials"})
		return
	}

	sessionToken := uuid.New().String()
	err = cache.SetEX(ctx, sessionToken, username, 60*60).Err()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Try again. Cannot set cookie"})
		return
	}

	c.SetCookie("session_token", sessionToken, 60*60, "/", "localhost:8000", false, false)
	c.JSON(http.StatusOK, gin.H{"message": "Login Successful", "user": username})
	return

}

//SignOut ...
func SignOut(c *gin.Context) {
	c.SetCookie("sessionToken", "", 0, "/", "localhost:8000", false, false)
	return
}
