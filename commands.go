package main

import (
	"fmt"
	"context"
	
	"github.com/karimOCB/blog_aggregator/internal/database"
)
type command struct {
	Name string
	Args []string
}

type commands struct {
	registry map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registry[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	value, ok := c.registry[cmd.Name]
	if !ok {
		return fmt.Errorf("the command: %q is not registered", cmd.Name)
	}
	return value(s, cmd)
}

func middlewareLoggedIn(handler func (s *state, cmd command, user database.User) error) func (*state, command) error {
	
	return func (s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

		if err != nil {
			return fmt.Errorf("could not retrieve user: %w", err)
		}

		return handler(s, cmd, user)
	}

}