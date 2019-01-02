package prompt

import (
	"testing"
)

func TestAdd(t *testing.T) {

	myCmd := Command{
		"ping",
		"just a test",
		func(args []string) {},
	}

	Add(myCmd)

	cmd := findCommand(myCmd.Cmd)
	if cmd.Action == nil {
		t.Error("Command not found")
	}
}

func TestAddCommand(t *testing.T) {

	cmdName := "foo"
	AddCommand(
		cmdName,
		"another test",
		func(args []string) {},
	)

	cmd := findCommand(cmdName)
	if cmd.Action == nil {
		t.Error("Command not found")
	}
}
