# go-commando [![GoDoc](https://godoc.org/github.com/binary-soup/go-commando?status.svg)](https://pkg.go.dev/github.com/binary-soup/go-commando)

The `go-command` module provides several types for managing multiple commands within the same command-line application. This approach to software architecture allows a single application to perform many tasks while ensuring the various commands stay modular and respect the _separation of concerns_ principal.

## Basic Usage

```go
func main() {
	ls := flag.Bool("ls", false, "list all commands")
	flag.Parse()

	runner := command.NewRunner(
		sample.NewHelloCommand(),
	)

	if *ls || len(os.Args) < 2 {
		runner.ListCommands()
		return
	}

	if err := runner.RunCommand(os.Args[1], os.Args[2:]); err != nil {
		style.BoldError.Print("ERROR: ")
		fmt.Println(err)
	}
}
```
