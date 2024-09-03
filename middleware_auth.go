package main

import (
	"fmt"
	"net/http"

	"github.com/iancenry/rss-feed-aggregator/handler"
	"github.com/iancenry/rss-feed-aggregator/internal/auth"
	"github.com/iancenry/rss-feed-aggregator/internal/database"
)


type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(h authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err  := auth.GetAPIKey(r.Header)

		if err != nil {
			handler.RespondWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err :=cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			handler.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		h(w, r, user)
	}
}	

