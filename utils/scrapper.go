package utils

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/iancenry/rss-feed-aggregator/internal/database"
)

func startScraping(db  *database.Queries, concurrency  int, timeBetweenRequest time.Duration) {
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
			go scrapeFeed(wg)
		}
		wg.Wait()

	}
	
}


func scrapeFeed(wg *sync.WaitGroup) {
	defer wg.Done()
}