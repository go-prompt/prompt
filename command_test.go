package prompt

import (
	"testing"
)

func TestAdd(t *testing.T) {

	my_cmd := Command{
		"ping",
		"just a test",
		func(args []string) {},
	}

	Add(my_cmd)

	if commands.Len() != 1 {
		t.Error("The command wasn't added")
	} else if commands[0].Cmd != "ping" {
		t.Error("Command not found")
	} else {
		// removes command added
		commands = commands[:commands.Len()-1]
	}
}

func TestAddCommand(t *testing.T) {

	AddCommand(
		"foo",
		"another test",
		func(args []string) {},
	)

	if commands.Len() != 1 {
		t.Error("The command wasn't added")
	} else if commands[0].Cmd != "foo" {
		t.Error("Command not found")
	} else {
		// removes command added
		commands = commands[:commands.Len()-1]
	}
}
