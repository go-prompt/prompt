package prompt_test

import (
	"fmt"
	"github.com/go-prompt/prompt"
)

func ExampleAdd() {
	cmd := prompt.Command{
		Cmd:  "ping",
		Desc: "just an example",
		Action: func() {
			fmt.Println("pong!")
		},
	}

	prompt.Add(cmd)

	<-prompt.Start() // if you don't need to lock, just remove the '<-' operator
}

func ExampleAddCommand() {
	prompt.AddCommand(
		"ping",
		"my custom command",
		func() {
			fmt.Println("pong!")
		},
	)

	<-prompt.Start() // if you don't need to lock, just remove the '<-' operator
}
