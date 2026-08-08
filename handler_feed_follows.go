package main

import (
	"context"
	"fmt"
	"time"

	"github.com/karimOCB/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("one argument is needed, a URL")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])

	if err != nil {
		return fmt.Errorf("couldn't retrieve feed by given URL: %w", err)
	}

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

func handlerFollowing(s *state, cmd command, user database.User) error {

	userFeedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("couldn't retrieve feed follows for current user: %w", err)
	}

	if len(userFeedFollows) == 0 {
		fmt.Println("no feeds followed")
	} else {
		for _, feedFollows := range userFeedFollows {
			fmt.Printf("Feed name: %s\n", feedFollows.FeedName)
		}
	}

	return nil
}


func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("one argument is needed, a URL")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	
	if err != nil {
		return fmt.Errorf("couldn't retrieve feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't delete followed feed: %w", err)
	}

	return nil
}