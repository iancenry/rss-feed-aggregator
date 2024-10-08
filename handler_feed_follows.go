package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/iancenry/rss-feed-aggregator/handler"
	"github.com/iancenry/rss-feed-aggregator/internal/database"
	"github.com/iancenry/rss-feed-aggregator/models"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)

	type paramters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := paramters{}

	err := decoder.Decode(&params)

	if err != nil {
		handler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		handler.RespondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v ", err))
		return
	}

	handler.RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedFollowToFeedFollow(feedFollow))

}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		handler.RespondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v ", err))
		return
	}

	handler.RespondWithJSON(w, http.StatusOK, models.DatabaseFeedFollowsToFeedFollows(feedFollows))

}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedId := chi.URLParam(r, "feedId")

	_, err := uuid.Parse(feedId)

	if err != nil {
		handler.RespondWithError(w, 400, "Invalid feed id")
		return
	}

	if err != nil {
		handler.RespondWithError(w, 400, fmt.Sprintf("Couldn't get feed follow: %v ", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID: uuid.MustParse(feedId),
	})
	if err != nil {
		handler.RespondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v ", err))
		return
	}
	handler.RespondWithJSON(w, http.StatusOK, struct{
		Message string `json:"message"`
	}{
		Message: "successfully deleted feed follow",
	})
}
