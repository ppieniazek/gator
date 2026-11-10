package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("couldn't wipe users table: %w", err)
	}
	fmt.Println("Database wiped!")
	return nil
}
