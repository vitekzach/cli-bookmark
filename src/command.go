package src

type Command struct {
	Command  string
	Category Category
	// CommandShell     Shell
	// ShellCommand     bool
	// LaunchExternally bool
	// SpecialHandler   string
}

type Category struct {
	Name string
}

type Shell struct {
	Name      string
	ShortName string
	Command   string
}
