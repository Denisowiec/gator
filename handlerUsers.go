package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	userlist, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving list of users from database: %v", err)
	}

	for _, u := range userlist {
		if s.cfg.CurrentUserName == u.Name {
			fmt.Printf("* %v (current)\n", u.Name)
		} else {
			fmt.Printf("* %v\n", u.Name)
		}
	}
	return nil
}
