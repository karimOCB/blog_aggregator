package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
	"strconv"

	"github.com/google/uuid"
	"github.com/karimOCB/blog_aggregator/internal/database"
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


func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.Args) != 0 {
		number, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("couldn't parse limit argument given: %w", err)
		}
		limit = int32(number)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: limit,
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Post: %+v\n", post)
	}

	return nil
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

	for _, item := range rssfeed.Channel.Item {
		pubDate := sql.NullTime{
			Valid: false,
		}
		
		parsedPubDate, err := time.Parse(time.RFC1123, item.PubDate)
		if err == nil {
			pubDate.Valid = true
			pubDate.Time = parsedPubDate
		} 
		
		err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{
				String: item.Title,
				Valid: item.Title != "",
			},
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid: item.Description != "",
			},
			PublishedAt: pubDate,
			FeedID: feed.ID,
		})
		
		if err != nil {
			if strings.Contains(err.Error(), "violates unique constraint") {
				continue
			} else {
				fmt.Printf("couldn't create post: %s", err)
			}
		}
	}

	return nil
}