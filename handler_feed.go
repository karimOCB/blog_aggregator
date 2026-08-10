package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/karimOCB/blog_aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {

	if len(cmd.Args) < 2 {
		return fmt.Errorf("two arguments are needed, name and url")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})

	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	fmt.Printf("Feed struct: %+v\n", feed)

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't follow the feed or already following: %w", err)
	}

	fmt.Printf("User: %s, Feed: %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not retrieve feeds, %w", err)
	}

	fmt.Printf("Feeds: %+v\n", feeds)

	return nil
}