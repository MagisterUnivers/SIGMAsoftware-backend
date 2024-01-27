package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Notes string    `json:"notes"`
}

var Users = make(map[uuid.UUID]User)
