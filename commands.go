package main

import (
	"fmt"
)

type command struct {
	name     string
	args     []string
	helpText string
}
type commands struct {
	cmdList map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdToRun, ok := c.cmdList[cmd.name]
	if !ok {
		return fmt.Errorf("command not found")
	}
	err := cmdToRun(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.cmdList == nil {
		c.cmdList = make(map[string]func(*state, command) error)
	}
	c.cmdList[name] = f
}
