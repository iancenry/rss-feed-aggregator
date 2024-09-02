package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/iancenry/rss-feed-aggregator/internal/database"
)

type User struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	ApiKey string `json:"api_key"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User {
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.CreatedAt,
		Name: dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}