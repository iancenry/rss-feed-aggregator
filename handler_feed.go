package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iancenry/rss-feed-aggregator/handler"
	"github.com/iancenry/rss-feed-aggregator/internal/database"
	"github.com/iancenry/rss-feed-aggregator/models"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)

	type paramters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}
	params := paramters{}

	err := decoder.Decode(&params)

	if err != nil {
		handler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Url:       params.URL,
		UserID:    user.ID,
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		handler.RespondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v ", err))
		return
	}

	handler.RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedToFeed(feed))

}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		handler.RespondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v ", err))
		return
	}

	handler.RespondWithJSON(w, http.StatusOK, models.DatabaseFeedsToFeeds(feeds))
}


// func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
// 	feeds, err := apiCfg.DB.GetFeed(r.Context(), user.ID)	

// 	if err != nil {
// 		handler.RespondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v ", err))
// 		return
// 	}

// 	handler.RespondWithJSON(w, http.StatusOK, feeds)
// }