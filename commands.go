package main

import (
	"fmt"

	"github.com/hyperflyguy/ThreatGator/internal/config"
)

type state struct {
	c *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if value, ok := c.registeredCommands[cmd.name]; ok {
		err := value(s, cmd)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Invalid command")
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
