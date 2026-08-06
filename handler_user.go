package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/karimOCB/blog_aggregator/internal/database"
)


func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("a username is needed to login")
	}

	username := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), username)

	if err != nil {
		return fmt.Errorf("You can't login to an account that does not exist. %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("Login successful")

	return nil
}


func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("a username is needed to register")
	}

	username := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})

	if err != nil {
		return err
	}

	s.cfg.SetUser(user.Name)
	fmt.Printf("The user was succesfully created. %+v\n", user)

	return nil
}


func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("Unsuccessful reset: %w", err)
	}

	fmt.Println("Successful reset")
	return nil
}


func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("could not retrieve users: %w", err)
	}

	currentLogged := s.cfg.CurrentUserName

	for _, user := range users {
		if user.Name == currentLogged {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}

	return nil
}


func handlerAgg(s *state, cmd command) error {
	rssfeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	
	if err != nil {
		return fmt.Errorf("could not fetch rssfeed: %w", err)
	}

	fmt.Printf("rssfeed struct: %+v\n", *rssfeed)

	return nil
}


func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not retrieve user: %w", err)
	}

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
		Url: url, 
		UserID: user.ID,
	})

	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	fmt.Printf("Feed struct: %+v\n", feed)

	return nil
}


func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not retrieve feeds, %w", err)
	}

	fmt.Printf("Feeds: %+v", feeds)

	return nil
}


func handlerFollow(s *state, cmd command) error {
	
	if len(cmd.Args) < 1 {
		return fmt.Errorf("one argument is needed, a URL")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	
	if err != nil {
		return fmt.Errorf("couldn't retrieve user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])

	if err != nil {
		return fmt.Errorf("couldn't retrieve feed by given URL: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't follow the feed or already following: %w", err)
	}

	fmt.Printf("User: %s, Feed: %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}


/*
func handlerFollowing(s *state, cmd command) error {

}
*/