package cli

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/codingchem/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: %s command expects a single arg. Found: %d", cmd.name, len(cmd.args))
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	err = s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("%s successfully logged in.\n", s.config.CurrentUser)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Error: Invalid number of arguments!")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == nil {
		return fmt.Errorf("Error: Username unavailable")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserName:  cmd.args[0],
	})
	if err != nil {
		return err
	}
	fmt.Printf("User successfully created!\n\tusername: %s\n\tid: %v\n", user.UserName, user.ID)
	s.cmds.run(s, command{
		name: "login",
		args: cmd.args,
	})
	return nil
}

func handlerReset(s *state, _ command) error {
	err := s.db.ResetUserTable(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("User table successfully reset!\n")
	return nil
}

func handlerListUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	if len(users) < 1 {
		fmt.Println("No users registered.")
	}
	for _, user := range users {
		if user.UserName == s.config.CurrentUser {
			fmt.Printf("* %s (current)\n", user.UserName)
		} else {
			fmt.Println("*", user.UserName)
		}
	}
	return nil
}
