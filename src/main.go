package main

import (
	"testapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.Auth(router)
	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}
}
