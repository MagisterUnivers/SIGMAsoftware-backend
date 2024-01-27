package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/google/uuid"
	"github.com/MagisterUnivers/SIGMAsoftware-backend/api"
	"github.com/MagisterUnivers/SIGMAsoftware-backend/models"
)

func main() {
	r := mux.NewRouter()

	idMap := map[string]uuid.UUID{
		"user1": uuid.New(),
		"user2": uuid.New(),
		"user3": uuid.New(),
	}

	user1ID := idMap["user1"]
	models.Users[user1ID] = models.User{
		ID:    user1ID,
		Name:  "John Wood",
		Email: "john@example.com",
		Notes: "Eg: Famous example guy",
	}

	user2ID := idMap["user2"]
	models.Users[user2ID] = models.User{
		ID:    user2ID,
		Name:  "Jane Foster",
		Email: "jane@example.com",
		Notes: "The one with the hammer",
	}

	user3ID := idMap["user3"]
	models.Users[user3ID] = models.User{
		ID:    user3ID,
		Name:  "Jackson Storm",
		Email: "jackson@example.com",
		Notes: "Not McQueen",
	}

	apiRoutes := r.PathPrefix("/api").Subrouter()
	api.UsersRoutes(apiRoutes.PathPrefix("/users").Subrouter())

	// Enable CORS
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(allowedOrigins, allowedMethods)(r))
}
