package prompt_test

import (
	"fmt"

	"github.com/go-prompt/prompt"
)

func ExampleAdd() {
	cmd := prompt.Command{
		Cmd:  "ping",
		Desc: "just an example",
		Action: func(args []string) {
			fmt.Println("pong! Arguments:", args)
		},
	}

	prompt.Add(cmd)

	<-prompt.Start() // if you don't need to lock, just remove the '<-' operator
}

func ExampleAddCommand() {
	prompt.AddCommand(
		"ping",
		"my custom command",
		func(args []string) {
			fmt.Println("pong! Arguments:", args)
		},
	)

	<-prompt.Start() // if you don't need to lock, just remove the '<-' operator
}
