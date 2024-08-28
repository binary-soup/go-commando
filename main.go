package main

import (
	"fmt"
	"local/command"
	"log"
	"os"
)

var commands []command.Command = []command.Command{}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments.")
	}

	if err := runCommand(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}

func runCommand(name string) error {
	for _, cmd := range commands {
		if cmd.GetName() == name {
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown command \"%s\"", name)
}
