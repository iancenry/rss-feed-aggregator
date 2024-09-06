package utils

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/iancenry/rss-feed-aggregator/internal/database"
)

func StartScraping(db  *database.Queries, concurrency  int, timeBetweenRequest time.Duration) {
	log.Printf("Starting scraping at %s \n", time.Now().Format(time.RFC3339))

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		// fetch x number of feeds from db, x being concurrency
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))

		if err != nil {
			log.Printf("Couldn't get feeds to fetch: %v", err)
			continue
		}

		// synchronization for multiple goroutines 
		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
	
}


func scrapeFeed(db *database.Queries ,wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("Couldn't mark feed as fetched: %v", err)
		return
	}
	rssFeed, err := UrlToFeed(feed.Url)

	if err != nil {
		log.Printf("Couldn't fetch feed: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		description := sql.NullString{}

		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubDate, err := time.Parse(time.RFC1123, item.PubDate)

		if err != nil {
			log.Printf("Couldn't parse date: %v with err  %v", item.PubDate ,  err)
			continue
		}


		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			FeedID: feed.ID,
			Url: item.Link,
			Title: item.Title,
			Description: description,
			PublishedAt: pubDate,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
		}
	}

}