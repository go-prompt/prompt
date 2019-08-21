package prompt

import (
	"fmt"
	"testing"
	"time"
)

func TestStart(t *testing.T) {

	Start()

	// waiting for commands (gorountine started)
	time.Sleep(time.Millisecond * 100)

	if commands.Len() == 0 {
		t.Error("No command found")
	}

	gcCmdName := "gc"
	cmd := findCommand(gcCmdName)
	if cmd.Cmd != gcCmdName {
		t.Errorf("GC command not found but '%s'\n", commands[0].Cmd)
	}
}

func TestRunCommand(t *testing.T) {
	cmd := &Command{
		Cmd:  "ping",
		Desc: "just an example",
		Action: func(args []string) {
			fmt.Println("pong! Arguments:", args)
		},
	}

	Add(*cmd)

	Start()

	// waiting for commands (gorountine started)
	time.Sleep(time.Millisecond * 100)

	cmdTarget := findCommand(cmd.Cmd)
	if cmdTarget.Cmd != cmd.Cmd {
		t.Errorf("%s command not found\n", cmd.Cmd)
	}

	cmdTarget.run([]string{})
}
