package main

import (
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Notes string `json:"notes"`
}

var users = make(map[uuid.UUID]User)

func main() {
	router := gin.Default()

	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.GET("/users/:id", getUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run(":8080")
}
