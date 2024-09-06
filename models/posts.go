package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/iancenry/rss-feed-aggregator/internal/database"
)

type  Post struct {
	ID          uuid.UUID `json:"id"`
	FeedID      uuid.UUID `json:"feed_id"`
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	Description *string `json:"description"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}


func DatabasePostToPost(dbPost database.Post) Post {
	var  description *string

	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		FeedID:      dbPost.FeedID,
		Url:         dbPost.Url,
		Title:       dbPost.Title,
		Description: description,
		Content:     dbPost.Content,
		PublishedAt: dbPost.PublishedAt,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
	}
}


func DatabasePostsToPosts(dbPosts []database.Post) []Post {
	var posts []Post

	for _, dbPost := range dbPosts {
		posts = append(posts, DatabasePostToPost(dbPost))
	}
	return posts
}