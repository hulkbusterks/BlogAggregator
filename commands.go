package main

import "errors"

type Command struct {
	name string
	args []string
}

type Commands struct {
	registred_commands map[string]func(*state, Command) error
}

func (c *Commands) register(name string, f func(*state, Command) error) {
	c.registred_commands[name] = f
}

func (c *Commands) run(s *state, cmd Command) error {
	f, ok := c.registred_commands[cmd.name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
