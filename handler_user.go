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
		return fmt.Errorf("You can't login to an account that does not exist. %s", err)
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
		return fmt.Errorf("Unsuccessful reset: %s", err)
	}

	fmt.Println("Successful reset")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("could not retrieve users: %s", err)
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
