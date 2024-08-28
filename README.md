# go-command [![GoDoc](https://godoc.org/github.com/binary-soup/go-command?status.svg)](https://pkg.go.dev/github.com/binary-soup/go-command)

The `go-command` module provides several types for managing multiple commands within the same command-line application. This approach to software architecture allows a single application to perform many tasks while ensuring the various commands stay modular and respect the _separation of concerns_ principal.

## Usage

The module consists mainly of three packages:
- `command`: this packages defines multiple types for defining, creating, and running commands
- `command/sample`: this package defines sample commands useful for learning how the module works
- `style`: this package provides semantic types and methods for changing the style (mode and colour) of console output

The following main definition demonstrates the basic usage of the module:

```go
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
```
