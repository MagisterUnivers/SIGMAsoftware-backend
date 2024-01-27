package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/MagisterUnivers/SIGMAsoftware-backend/api"
	"github.com/MagisterUnivers/SIGMAsoftware-backend/models"
)

func main() {
	router := gin.Default()

	idMap := map[string]uuid.UUID{
		"user1": uuid.New(),
		"user2": uuid.New(),
		"user3": uuid.New(),
	}

	user1ID := idMap["user1"]
	models.Users[user1ID] = models.User{ 
		ID: user1ID,
		Name: "John Wood", 
		Email: "john@example.com", 
		Notes: "Eg: Famous example guy",
	 }

	user2ID := idMap["user2"]
  models.Users[user2ID] = models.User{ 
		ID: user2ID,
		Name: "Jane Foster",
		Email: "jane@example.com", 
		Notes: "The one with the hammer",
		}

	user3ID := idMap["user3"]
  models.Users[user3ID] = models.User{
		ID: user3ID, 
		Name: "Jackson Storm", 
		Email: "jackson@example.com", 
		Notes: "Not McQueen",
	}

	apiRoutes := router.Group("/api")
	{
		api.UsersRoutes(apiRoutes.Group("/users"))
	}

	router.Run(":8080")
}
