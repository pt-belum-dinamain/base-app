package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pt-belum-dinamain/base-app/internal/base_app/controllers"
)

func main() {
	router := gin.Default()
	user := router.Group("/api/user")
	{
		user.POST("/", controllers.CreateUser)
		user.GET("/", controllers.GetUser)
		user.GET("/:id", controllers.GetUserById)
		user.PUT("/", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	role := router.Group("/api/role")
	{
		role.POST("/", controllers.CreateRole)
		role.GET("/", controllers.GetRole)
		role.GET("/:id", controllers.GetRoleById)
		role.PUT("/", controllers.UpdateRole)
		role.DELETE("/:id", controllers.DeleteRole)
	}
	router.Run()
}
