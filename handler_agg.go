package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	tempURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), tempURL)
	if err != nil {
		return fmt.Errorf("error fetching '%s': %w", tempURL, err)
	}
	fmt.Println(feed)
	return nil
}
