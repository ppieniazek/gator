package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ppieniazek/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	usage := fmt.Errorf("usage: %s <optional number starting from 2>", cmd.Name)
	postsLimit := 2
	if len(cmd.Args) == 1 {
		limit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return usage
		}
		postsLimit = limit
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(postsLimit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}
	if len(posts) == 0 {
		fmt.Println("There are no fetched posts yet!")
		return nil
	}
	for _, post := range posts {
		printPost(post)
	}
	return nil
}

func printPost(post database.GetPostsForUserRow) {
	fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
	fmt.Printf("--- %s ---\n", post.Title)
	fmt.Printf("    %v\n", post.Description.String)
	fmt.Printf("Link: %s\n", post.Url)
	fmt.Println("=====================================")
}
