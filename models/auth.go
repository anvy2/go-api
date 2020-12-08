package models

import (
	"testapi/controller"

	"github.com/gin-gonic/gin"
)

// Auth -> Adds authentication endpoints
func Auth(r *gin.Engine) {
	// controller.Connect()
	r.POST("/signUp", controller.SignUp)
	r.POST("/login", controller.SignIn)
	r.POST("/logout", controller.SignOut)
}
