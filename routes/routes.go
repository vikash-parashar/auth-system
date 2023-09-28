// routes/routes.go
package routes

import (
	"auth-system/controllers"
	"auth-system/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Private routes (require authentication)
	private := r.Group("/private")
	private.Use(utils.AuthMiddleware("user"))
	{
		private.GET("/user", controllers.GetUser)
	}

	return r
}
