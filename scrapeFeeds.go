package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Denisowiec/gator/internal/database"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	params := database.MarkFeedFetchedParams{
		ID:            feed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}
	_, err = s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return err
	}

	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("Channel name: %v\n", fetchedFeed.Channel.Title)

	for _, item := range fetchedFeed.Channel.Item {
		fmt.Printf("Article title: %v\n", item.Title)
	}

	return nil
}
