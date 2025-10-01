package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Denisowiec/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("not enough arguments")
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]

	createFeedParams := database.CreateFeedParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), createFeedParams)
	if err != nil {
		return fmt.Errorf("error adding feed: %v", err)
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}
	fmt.Printf("Added feed %v\n", feedName)

	_, err = s.db.CreateFeedFollow(context.Background(), createFeedFollowParams)
	if err != nil {
		return fmt.Errorf("error adding a follow: %v", err)
	}
	return nil
}
