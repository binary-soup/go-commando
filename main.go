package main

import (
	"flag"
	"local/command"
	"local/command/sample"
	"log"
	"os"
)

func main() {
	ls := flag.Bool("ls", false, "list all commands")
	flag.Parse()

	runner := command.NewRunner(sample.NewHelloCommand())

	if *ls || len(os.Args) < 2 {
		runner.ListCommands()
		return
	}

	if err := runner.RunCommand(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}
