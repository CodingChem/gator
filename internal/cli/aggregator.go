package cli

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/codingchem/gator/internal/database"
	"github.com/codingchem/gator/internal/rss"
	"github.com/google/uuid"
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
	for _, post := range feed.Channel.Item {

		s.db.CreatePost(context.Background(),database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			Title: post.Title,
			Url: post.Link,
			Description: sql.NullString{String:post.Description, Valid: post.Description != ""},
			//TODO: Implement the nulltime parser
			PublishedAt: sql.NullTime{Valid: false},
			FeedID: next.ID,

		})
	}
}
