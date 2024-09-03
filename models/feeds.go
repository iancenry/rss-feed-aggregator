package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/iancenry/rss-feed-aggregator/internal/database"
)

type Feed struct {
	ID uuid.UUID `json:"id"`
	Url string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DatabaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed {
		ID: dbFeed.ID,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
		Name: dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func DatabaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	var feeds []Feed
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, DatabaseFeedToFeed(dbFeed))
	}
	return feeds
}

