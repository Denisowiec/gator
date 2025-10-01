package main

import (
	"context"
	"fmt"

	"github.com/Denisowiec/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, item := range feeds {
		feed, err := s.db.GetFeedByID(context.Background(), item.FeedID)
		if err != nil {
			return err
		}
		fmt.Printf("Feed name: %s\n", feed.Name)
	}
	return nil
}
