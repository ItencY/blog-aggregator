package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	runCommand, ok := c.commands[cmd.name]
	if !ok {
		return errors.New("unknown command: " + cmd.name)
	}
	return runCommand(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
