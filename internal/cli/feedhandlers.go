package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/codingchem/gator/internal/database"
	"github.com/codingchem/gator/internal/rss"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: agg takes a single argument!")
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("Collecting feeds every", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Error: Addfeed requires 2 arguments!")
	}

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.args[0],
			Url:       cmd.args[1],
			UserID:    user.ID,
		})
	if err != nil {
		return err
	}
	_, err = s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID: uuid.New(),	
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Error: feeds takes no arguments!")
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("\nTitle: %s\nURL: %s\nCreated by: %s\n", feed.Name, feed.Url, feed.UserName)
	}
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: follow takes a single argument!")
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return nil
	}

	inserted_feed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return nil
	}
	fmt.Printf("User: %s\n follows: %s\n",inserted_feed.UserName, inserted_feed.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Error: following takes no arguments!")
	}
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	fmt.Printf("%s follows the following feeds:\n",user.UserName)
	for _, feed := range feeds {
		fmt.Printf("\t- %s\n",feed.FeedName)
	}
	return nil
}

func handlerUnFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: unfollow takes a single argument!")
	}
	feed, err := s.db.GetFeedByUrl(context.Background(),cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(),database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	return err
}

func printFeed(feed *rss.RSSFeed) {
	fmt.Printf("Feed: %s\nDescription: %s\n", feed.Channel.Title, feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("- %s\n\tLink: %s\n", item.Title, item.Link)
	}
}
