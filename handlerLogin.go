package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no arguments given")
	}
	newuser := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), newuser)
	if err != nil {
		return fmt.Errorf("error loging into databas: %v", err)
	}

	if err := s.cfg.SetUser(newuser); err != nil {
		return err
	}
	fmt.Printf("New user %v set\n", newuser)
	return nil
}
