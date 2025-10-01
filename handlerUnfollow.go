package main

import (
	"context"
	"fmt"

	"github.com/Denisowiec/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("not enough arguments")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, err = s.db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}
	return nil
}
