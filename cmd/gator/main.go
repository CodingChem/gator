package main

import (
	"fmt"
	"log"

	"github.com/codingchem/gator/internal/config"
)

func main() {
	globalConfig, err := config.Read()
	if err != nil {
		log.Panicf(err.Error())
	}
	fmt.Printf("url: %s\nuser: %s\n", globalConfig.DB_CON_STRING, globalConfig.CurrentUser)
	err = globalConfig.SetUser("vegard")
	if err != nil {
		log.Panicf(err.Error())
	}
	fmt.Println("Loading new user")
	globalConfig, err = config.Read()
	if err != nil {
		log.Panicf(err.Error())
	}
	fmt.Printf("url: %s\nuser: %s\n", globalConfig.DB_CON_STRING, globalConfig.CurrentUser)
}
