package main

import (
	"flag"
	"os"

	"github.com/binary-soup/go-command/alert"
	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/command/sample"
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
		alert.Print(err)
	}
}
