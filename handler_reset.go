package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("Unsuccessful reset: %w", err)
	}

	fmt.Println("Successful reset")
	return nil
}
