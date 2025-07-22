/*
# Sample Command

	type FooCommand struct {
		command.CommandBase
		flags *fooFlags
	}

	type fooFlags struct {
	}

	func (f *fooFlags) Set(flags *flag.FlagSet) {
	}

	func NewFooCommand() FooCommand {
		flags := new(fooFlags)

		return FooCommand{
			CommandBase: command.NewCommandBase("", "", flags),
			flags:       flags,
		}
	}

	func (cmd FooCommand) Run() error {
		return nil
	}
*/
package command
