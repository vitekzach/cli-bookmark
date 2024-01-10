package src

type Command struct {
	Command          string
	Category         string
	CommandShell     Shell
	ShellCommand     bool
	LaunchExternally bool
	SpecialHandler   string
}

type Shell struct {
	Name      string
	ShortName string
	Command   string
}
