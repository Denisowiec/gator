package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting the database: %v", err)
	}

	fmt.Printf("database reset successfully\n")
	return nil
}
