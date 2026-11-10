package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ppieniazek/gator/internal/database"
)

func handlerListFollowedFeeds(s *state, cmd command) error {
	currentUser, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feed_follows, err := s.db.ListFeedFollowsByUserID(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("couldn't list feed_follows: %w", err)
	}
	if len(feed_follows) == 0 {
		fmt.Println("You are not following any feeds.")
		return nil
	}

	fmt.Printf("Followed feeds for user %s:\n", currentUser.Name)
	for _, feed_follow := range feed_follows {
		fmt.Printf(" * Name:	%v\n", feed_follow.FeedName)
	}
	return nil
}

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	currentUser, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("couldn't create feed_follow: %w", err)
	}
	fmt.Println("Feed followed:")
	fmt.Printf(" * Name:	%v\n", feedFollow.FeedName)
	fmt.Printf(" * User:	%v\n", feedFollow.UserName)
	return nil
}
