package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Denisowiec/gator/internal/database"
	"github.com/lib/pq"
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

	for _, item := range fetchedFeed.Channel.Item {
		// We need to convert the publishing date:
		pubdate, err := time.Parse(time.RFC822, item.PubDate)
		if err != nil {
			pubdate = time.Now()
		}
		cp_params := database.CreatePostParams{
			CreatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubdate,
			FeedID:      feed.ID,
		}
		_, err = s.db.CreatePost(context.Background(), cp_params)
		if err != nil {
			// if there's a duplicate url in the databse, we can ignore it.
			if err, ok := err.(*pq.Error); ok && err.Code == "23505" {
				continue
			} else {
				// else: we print all the details to improve the error handling in the future
				fmt.Println("Severity:", err.Severity)
				fmt.Println("Code:", err.Code)
				fmt.Println("Message:", err.Message)
				fmt.Println("Detail:", err.Detail)
				fmt.Println("Hint:", err.Hint)
				fmt.Println("Position:", err.Position)
				fmt.Println("InternalPosition:", err.InternalPosition)
				fmt.Println("Where:", err.Where)
				fmt.Println("Schema:", err.Schema)
				fmt.Println("Table:", err.Table)
				fmt.Println("Column:", err.Column)
				fmt.Println("DataTypeName:", err.DataTypeName)
				fmt.Println("Constraint:", err.Constraint)
				fmt.Println("File:", err.File)
				fmt.Println("Line:", err.Line)
				fmt.Println("Routine:", err.Routine)
			}
		}
	}

	return nil
}
