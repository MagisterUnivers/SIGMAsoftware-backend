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
