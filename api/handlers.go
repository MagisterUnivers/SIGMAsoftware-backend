package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/MagisterUnivers/SIGMAsoftware-backend/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uuid.New()
	models.Users[user.ID] = user
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	usersList := make([]models.User, 0, len(models.Users))
	for _, user := range models.Users {
		usersList = append(usersList, user)
	}
	c.JSON(http.StatusOK, usersList)
}

func GetUser(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, exists := models.Users[userID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, exists := models.Users[userID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Users[userID] = user
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	_, exists := models.Users[userID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	delete(models.Users, userID)
	c.JSON(http.StatusNoContent, nil)
}
