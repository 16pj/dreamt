package cmdline

import "fmt"

// create an entity that handles commandline commands
type Cmd struct {
	// create a map of commands
	commands map[string]func([]string) error
}

// create a constructor for the Cmd struct
func NewCmd() *Cmd {
	return &Cmd{
		commands: make(map[string]func([]string) error),
	}
}

// create a method that adds a command to the Cmd struct
func (cmd *Cmd) Add(name string, handler func([]string) error) {
	cmd.commands[name] = handler
}

// create a method that executes a command
func (cmd *Cmd) Execute(name string, args []string) error {
	if handler, ok := cmd.commands[name]; ok {
		return handler(args)
	}

	return nil
}

// create a method that prints the help for all commands
func (cmd *Cmd) Help() {
	for name := range cmd.commands {
		fmt.Printf("%s\n", name)
	}
}

// create a method that prints the help for a specific command
func (cmd *Cmd) HelpCmd(name string) {
	if handler, ok := cmd.commands[name]; ok {
		handler(nil)
	}
}

// create a method that prints the help for all commands
func (cmd *Cmd) HelpAll() {
	for name, handler := range cmd.commands {
		handler(nil)
		fmt.Printf("%s\n", name)
	}
}

// create a method that prints the help for all commands
func (cmd *Cmd) HelpAllCmd(name string) {
	if handler, ok := cmd.commands[name]; ok {
		handler(nil)
	}
}

// create a method that prints the help for all commands
func (cmd *Cmd) HelpAllCmds(name string) {
	if handler, ok := cmd.commands[name]; ok {
		handler(nil)
	}
}

// create a method that prints the help for all commands
func (cmd *Cmd) HelpAllCommands(name string) {
	if handler, ok := cmd.commands[name]; ok {
		handler(nil)
	}
}
