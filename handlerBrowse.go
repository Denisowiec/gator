package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Denisowiec/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	// limit is an SQL parameter, so we need to convert it to the correct format
	var limit int32
	limit = 2
	if len(cmd.args) > 0 {
		arglimit, err := strconv.Atoi(cmd.args[0])
		if err == nil {
			limit = int32(arglimit)
		}
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	getpostsParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}
	posts, err := s.db.GetPostsForUser(context.Background(), getpostsParams)
	if err != nil {
		return err
	}

	for _, post := range posts {
		feed, err := s.db.GetFeedByID(context.Background(), post.FeedID)
		if err != nil {
			return err
		}
		fmt.Printf("Feed: %v\n", feed.Name)
		fmt.Printf("Title: %v\nPublished: %v\nDescription:%v\n\n", post.Title, post.PublishedAt, post.Description)
	}

	return nil
}
