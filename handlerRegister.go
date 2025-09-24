package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Denisowiec/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no arguments given")
	}

	dbArgs := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}
	newuser, err := s.db.CreateUser(context.Background(), dbArgs)

	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	s.cfg.SetUser(newuser.Name)

	fmt.Printf("User %v was created\n", newuser.Name)
	return nil
}
