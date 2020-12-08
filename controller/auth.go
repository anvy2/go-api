package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testapi/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client *mongo.Client   = db.GetClient()
	cache  *redis.Client   = db.InitRedis()
	ctx    context.Context = context.Background()
)

type person struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func serialize(s []byte) person {
	var p person
	err := json.Unmarshal(s, &p)
	if err != nil {
		panic(err)
	}
	return p
}

//SignUp ...
func SignUp(c *gin.Context) {
	collection := client.Database("testapi").Collection("user_table")

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := buf[0:num]
	u := serialize(reqBody)
	fmt.Println(u)
	username := u.Username
	password := u.Password

	user := db.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: password,
	}
	var e db.User
	err := collection.FindOne(context.TODO(), bson.D{{}}).Decode(&e)
	if err == mongo.ErrNoDocuments {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "database error"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "user created"})
		return
	}
	c.JSON(http.StatusAlreadyReported, gin.H{"message": "User already exists!", "username": username, "password": password})
	return
}

//SignIn ...
func SignIn(c *gin.Context) {
	collection := client.Database("testapi").Collection("user_table")

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := buf[0:num]
	u := serialize(reqBody)
	// fmt.Println(u)
	username := u.Username
	password := u.Password
	var user db.User
	filter := bson.M{
		"username": username,
	}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusOK, gin.H{"message": "No user found", "filter": filter, "user": user, "username": username})
		return
	}

	if user.Password != password {
		fmt.Println(filter)
		c.JSON(http.StatusOK, gin.H{"message": "Invalid Credetials"})
		return
	}

	sessionToken := uuid.New().String()
	err = cache.SetEX(ctx, sessionToken, username, 60*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Try again. Cannot set cookie"})
		return
	}

	c.SetCookie("session_token", sessionToken, 60, "/login", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"message": "Login Successful", "user": username})
	return

}

//SignOut ...
func SignOut(c *gin.Context) {
	c.SetCookie("sessionToken", "", 0, "/login", "localhost", false, false)
	return
}
