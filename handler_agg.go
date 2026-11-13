package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ppieniazek/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time>", cmd.Name)
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse the time: %w", err)
	}
	const fastestTimeBetweenRequests = time.Second * 5
	if timeBetweenRequests < fastestTimeBetweenRequests {
		fmt.Println("You cannot set minimal time less than:", fastestTimeBetweenRequests)
		return nil
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("Couldn't get the next feed: %v", err)
		return
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed as fetched: %v", err)
		return
	}
	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't fetch the feed: %v", err)
		return
	}
	fmt.Printf("Fetched '%v' feed!\n", fetchedFeed.Channel.Title)
	for _, post := range fetchedFeed.Channel.Item {
		savePost(s, post, feed)
	}
}

func savePost(s *state, post RSSItem, feed database.Feed) {
	err := s.db.CreatePost(context.Background(), database.CreatePostParams{
		Title: post.Title,
		Url:   post.Link,
		Description: sql.NullString{
			String: post.Description,
			Valid:  post.Description != "",
		},
		PublishedAt: parsePubDate(post.PubDate),
		FeedID:      feed.ID,
	})
	if err != nil && !strings.Contains(err.Error(), "23505") {
		log.Printf("Couldn't create the post: %v", err)
	}
}

func parsePubDate(s string) sql.NullTime {
	if s == "" {
		return sql.NullTime{}
	}

	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC3339,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
	}

	return sql.NullTime{}
}
