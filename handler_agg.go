package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("a time between request argument is needed")
	}

	timeBetweenReq, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse given time: %w", err)
	}

	fmt.Printf("Collecting feeds every: %s\n", timeBetweenReq)

	ticker := time.NewTicker(timeBetweenReq)
	for ; ; <-ticker.C {
		
		scrapeFeeds(s)
	}

}


func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	
	if err != nil {
		return fmt.Errorf("couldn't fetch next feed: %w", err)
	} 

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}

	rssfeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("could not fetch rssfeed: %w", err)
	}

	fmt.Println("\nNext Feed: ")

	for _, item := range rssfeed.Channel.Item {
		fmt.Printf("item title: %s \n", item.Title)
	}

	return nil
}