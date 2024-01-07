package src

type Command struct {
	Command          string
	Category         string
	CommandShell     Shell
	ShellCommand     bool
	LaunchExternally bool
}

type Shell struct {
	Name    string
	Command string
}
