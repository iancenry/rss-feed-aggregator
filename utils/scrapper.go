package utils

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/iancenry/rss-feed-aggregator/internal/database"
)

func StartScraping(db  *database.Queries, concurrency  int, timeBetweenRequest time.Duration) {
	log.Printf("Starting scraping at %s", time.Now().Format(time.RFC3339))
	log.Printf("Scraping on  %v goroutines at %s duration", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		log.Printf("Scraping at %s", time.Now().Format(time.RFC3339))
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
		log.Println("Found post", item.Title + "on feed " + feed.Name)
	}

	_, err = db.CreatePost(context.Background(), database.CreatePostParams{
    FeedID      feed.ID,
    Url         rssFeed.Channel.Items[0].Link,
    Title       rssFeed.Channel.Items[0].Title,
    Description rssFeed.Channel.Items[0].Description,
    Content     rssFeed.Channel.Items[0].Content,
    PublishedAt rssFeed.Channel.Items[0].PubDate
	})
}