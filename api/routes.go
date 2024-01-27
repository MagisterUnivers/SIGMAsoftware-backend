package api

import (
	"github.com/gin-gonic/gin"
)

func UsersRoutes(group *gin.RouterGroup) {
	group.POST("/", CreateUser)
	group.GET("/", GetUsers)
	group.GET("/:id", GetUser)
	group.PUT("/:id", UpdateUser)
	group.DELETE("/:id", DeleteUser)
}
