package cli

import (
	"context"
	"log"
	"time"

	"github.com/codingchem/gator/internal/database"
	"github.com/codingchem/gator/internal/rss"
)

func scrapeFeeds(s *state){
	next, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Fatal("Fatal Error: %w\n", err)
	}
	err = s.db.MarkFeedFetched(context.Background(),database.MarkFeedFetchedParams{
		ID: next.ID,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Fatal("Fatal Error: %w\n", err)
	}
	feed, err := rss.FetchFeed(context.Background(), next.Url)
	if err != nil {
		log.Fatal("Fatal Error: %w\n", err)
	}
	printFeed(feed)
}
