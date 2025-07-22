package command

// The interface definition of a command.
// Any type that implements the interface can be used by the runner.
type Command interface {
	// Returns the command's name; essentially it's identifier.
	GetName() string

	// Prints usage information.
	PrintUsage()

	// Submits the arguments for execution. If already submitted, the new arguments should override the old ones.
	SubmitArgs(args []string) error

	// Runs the command. If arguments are desired, SubmitArgs should be called prior.
	Run() error
}
