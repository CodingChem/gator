package cli

import (
	"fmt"
)

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
	err = c.register("register", "\tRegister a new user\n\t`Usage: gator register $user`\n", handlerRegister)
	if err != nil {
		return commands{}, err
	}
	err = c.register("reset", "\tReset users table Warning: Dangerous!\n\tUsage: `gator reset`\n", handlerReset)
	if err != nil {
		return commands{}, err
	}
	err = c.register("users", "\tList registered users\n\tUsage: `gator list`\n", handlerListUsers)
	if err != nil {
		return commands{}, err
	}
	err = c.register("agg", "\tFetch and print a feed\n\tUsage: `gator agg`\n", handlerAgg)
	if err != nil {
		return commands{}, err
	}
	err = c.register("addfeed", "\tAdd a feed\n\tUsage: `gator addfeed $name $url`\n", handlerAddFeed)
	if err != nil {
		return commands{}, err
	}
	err = c.register("feeds", "\tList all feeds\n\tUsage: `gator feeds`\n", handlerListFeeds)
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

func handlerHelp(s *state, _ command) error {
	fmt.Printf("Welcome to gator!\n\nUsage:\n")
	for _, handler := range s.cmds.commandMap {
		fmt.Println()
		fmt.Printf("Command: %s\nDescription:\n%s\n", handler.name, handler.description)
	}
	return nil
}
