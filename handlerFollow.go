package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Denisowiec/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("error: not enough arguments")
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error accessing user database: %v", err)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("errof accessing feed database: %v", err)
	}

	params := database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedfollow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating a feedfollow pair: %v", err)
	}

	fmt.Printf("Feed name: %s\nUser name: %s\n", feedfollow.FeedName, feedfollow.UserName)
	return nil
}
