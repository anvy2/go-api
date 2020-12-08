package main

import (
	"testapi/db"
	"testapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	client, ctx, _ := db.Connect("localhost:27017")
	defer client.Disconnect(ctx)
	router := gin.Default()
	models.Auth(router)
	err := router.Run(":8000")
	panic(err)
}
