package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

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
