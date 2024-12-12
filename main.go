package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codingchem/gator/internal/config"
)

func main() {
	globalState, err := initState()
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Incorrect usage, consult 'help' for help")
		os.Exit(1)
	}
	err = globalState.cmds.run(&globalState, command{
		name: args[1],
		args: args[2:],
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func initState() (state, error) {
	globalConfig, err := config.Read()
	if err != nil {
		return state{}, err
	}
	cmds, err := NewCommands()
	if err != nil {
		return state{}, err
	}
	return state{
		config: &globalConfig,
		cmds:   &cmds,
	}, nil
}
