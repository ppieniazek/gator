package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	if err := s.cfg.SetUser(username); err != nil {
		return err
	}
	fmt.Println("User switched successfully!")
	return nil
}
