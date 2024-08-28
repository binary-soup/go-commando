package command

type Command interface {
	GetName() string
	PrintUsage()
	Run() error
}
