package main

import "fmt"

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