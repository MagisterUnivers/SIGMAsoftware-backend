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

	users[uuid.New()] = User{ Name: "John Wood", Email: "john@example.com", Notes: "Eg: Famous example guy" }
  users[uuid.New()] = User{ Name: "Jane Foster", Email: "jane@example.com", Notes: "The one with the hammer" }
  users[uuid.New()] = User{ Name: "Jackson Storm", Email: "jackson@example.com", Notes: "Not McQueen" }

	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.GET("/users/:id", getUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run(":8080")
}

func createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uuid.New()
	users[user.ID] = user
	c.JSON(http.StatusCreated, user)
}
