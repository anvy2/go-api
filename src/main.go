package main

import (
	"testapi/db"
	"testapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect("mongodb://root:password@localhost:27017")
	db.InitRedis()
	client := db.Client
	ctx := db.Ctx
	defer client.Disconnect(ctx)
	router := gin.Default()
	models.Auth(router)
	err := router.Run(":8000")
	panic(err)
}
