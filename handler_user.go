package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ppieniazek/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]
	if s.cfg.CurrentUserName == username {
		fmt.Println("already logged in")
		return nil
	}
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}
	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User switched!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	newUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      username,
	}
	newUser, err := s.db.CreateUser(context.Background(), newUserParams)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	if err := s.cfg.SetUser(newUser.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User created:")
	printUser(newUser)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}
	if len(users) == 0 {
		fmt.Println("There are no users yet!")
		return nil
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Println(" * ", user.Name, "(current)")
			continue
		}
		fmt.Println(" * ", user.Name)
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:		%v\n", user.ID)
	fmt.Printf(" * Name:	%v\n", user.Name)
}
