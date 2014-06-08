package prompt

import (
	"fmt"
)

type Command struct {
	Cmd    string
	Desc   string
	Action func()
}

func (c *Command) Run() bool {

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

	fmt.Printf("%s: command not found.\n", c.Cmd)
	return false
}

func Parse(txt string) *Command {

	for _, cmd := range commands {
		if cmd.Cmd == txt {
			return &cmd
		}
	}

	return &Command{Cmd: txt, Action: nil}
}
