package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	/*url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error fetching the feed: %v", err)
	}
	fmt.Println(feed)*/

	if len(cmd.args) < 1 {
		return fmt.Errorf("not enough arguments")
	}

	interval, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", interval)

	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {
		fmt.Println("Scraping...")
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

	//return nil
}
