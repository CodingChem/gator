package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codingchem/gator/internal/cli"
	_ "github.com/lib/pq"
)

func main() {
	globalState, err := cli.NewState()
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Incorrect usage, consult 'help' for help")
		os.Exit(1)
	}
	err = globalState.Run(args[1], args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
