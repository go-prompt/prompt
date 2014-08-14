package prompt

import (
	"fmt"
	"sort"
)

// Command is a representation of a valid command.
type Command struct {
	Cmd    string
	Desc   string
	Action func()
}

func (c *Command) run() bool {

	if c.Cmd == "quit" {
		return true
	} else if c.Cmd == "" {
		return false
	}

	if c.Action != nil {

		defer func() { // handling panic
			if r := recover(); r != nil {
				fmt.Println("Ops! ["+c.Cmd+"] -", r)
			}
		}()

		c.Action()
		return false
	}

	fmt.Printf("%s: command not found. Try 'help'.\n", c.Cmd)
	return false
}

type commandList []Command

func (c commandList) Len() int           { return len(c) }
func (c commandList) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c commandList) Less(i, j int) bool { return c[i].Cmd < c[j].Cmd }

var commands commandList

func findCommand(txt string) *Command {

	for _, cmd := range commands {
		if cmd.Cmd == txt {
			return &cmd
		}
	}

	return &Command{Cmd: txt, Action: nil}
}

// Add provides a new Command.
func Add(c Command) {

	commands = append(commands, c)
	sort.Sort(commands)
}

// AddCommand constructs a new Command and provides it.
func AddCommand(name string, desc string, action func()) {

	Add(Command{name, desc, action})
}
