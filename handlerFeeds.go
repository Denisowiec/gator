package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error accessing feed database: %v", err)
	}

	for _, item := range feeds {
		user, err := s.db.GetUserByUUID(context.Background(), item.UserID)
		if err != nil {
			return fmt.Errorf("error while referencing user id: %v", err)
		}
		fmt.Printf("Feed name: %s\n", item.Name)
		fmt.Printf("Feed url: %s\n", item.Url)
		fmt.Printf("User: %s\n", user.Name)
	}
	return nil
}
