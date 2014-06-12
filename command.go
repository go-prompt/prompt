package prompt

import (
	"fmt"
	"sort"
)

type CommandList []Command

func (c CommandList) Len() int           { return len(c) }
func (a CommandList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CommandList) Less(i, j int) bool { return a[i].Cmd < a[j].Cmd }

var commands CommandList

type Command struct {
	Cmd    string
	Desc   string
	Action func()
}

func Add(c Command) {

	commands = append(commands, c)
	sort.Sort(commands)
}

func AddCommand(name string, desc string, action func()) {

	Add(Command{name, desc, action})
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
				fmt.Println(r)
			}
		}()

		c.Action()
		return false
	}

	fmt.Printf("%s: command not found. Try 'help'.\n", c.Cmd)
	return false
}

func parse(txt string) *Command {

	for _, cmd := range commands {
		if cmd.Cmd == txt {
			return &cmd
		}
	}

	return &Command{Cmd: txt, Action: nil}
}
