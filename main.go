package main

import (
	"flag"
	"fmt"
	"local/command"
	"local/command/sample"
	"local/style"
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

	if err := runner.RunCommand(os.Args[1], os.Args[2:]); err != nil {
		style.BoldError.Print("ERROR: ")
		fmt.Println(err)
	}
}
