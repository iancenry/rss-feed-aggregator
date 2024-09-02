package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iancenry/rss-feed-aggregator/handler"
	"github.com/iancenry/rss-feed-aggregator/internal/auth"
	"github.com/iancenry/rss-feed-aggregator/internal/database"
	"github.com/iancenry/rss-feed-aggregator/models"
)

// since  we want to pass another argument ie apiConfig
// we're chaning this func to a method with a receiver apiCfg which is a pointer to an apiConfig
// so now we can access additional data stored on the struct itself
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		handler.RespondWithError(w, http.StatusCreated, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		handler.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v ", err))
		return
	}

	handler.RespondWithJSON(w, 200, models.DatabaseUserToUser(user))
} 

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err  := auth.GetAPIKey(r.Header)

	if err != nil {
		handler.RespondWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		handler.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	handler.RespondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))

}