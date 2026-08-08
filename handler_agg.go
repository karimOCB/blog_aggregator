package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	rssfeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return fmt.Errorf("could not fetch rssfeed: %w", err)
	}

	fmt.Printf("rssfeed struct: %+v\n", *rssfeed)

	return nil
}
