package main

import (
	"fmt"

	"github.com/codingchem/gator/internal/config"
)

type state struct {
	config *config.Config
	cmds   *commands
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]handler
}

type handler struct {
	cmd         func(*state, command) error
	name        string
	description string
}

func NewCommands() (commands, error) {
	var c commands
	c.commandMap = make(map[string]handler)
	err := c.register("login", "\tLogin as user\n\t`Usage: gator login $user`\n", handlerLogin)
	if err != nil {
		return commands{}, err
	}
	err = c.register("help", "\tDisplays a friendly help message\n\t`Usage: gator help`\n", handlerHelp)
	if err != nil {
		return commands{}, err
	}
	return c, nil
}

func (c *commands) register(name string, description string, f func(*state, command) error) error {
	if _, exists := c.commandMap[name]; exists {
		return fmt.Errorf("Error: Command with name: %s already in commands!", name)
	}
	c.commandMap[name] = handler{name: name, description: description, cmd: f}
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.commandMap[cmd.name]
	if !exists {
		return fmt.Errorf("Error: No function '%s' in commands!", cmd.name)
	}
	return handler.cmd(s, cmd)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: %s command expects a single arg. Found: %d", cmd.name, len(cmd.args))
	}
	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("%s successfully logged in.\n", s.config.CurrentUser)
	return nil
}

func handlerHelp(s *state, _ command) error {
	fmt.Printf("Welcome to gator!\n\nUsage:\n")
	for _, handler := range s.cmds.commandMap {
		fmt.Println()
		fmt.Printf("Command: %s\nDescription:\n%s", handler.name, handler.description)
	}
	return nil
}
