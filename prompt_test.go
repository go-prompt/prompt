package prompt

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {

	Start()

	// waiting for commands (gorountine started)
	time.Sleep(time.Millisecond * 100)

	if commands.Len() == 0 {
		t.Error("No command found")
	} else if commands[0].Cmd != "gc" {
		t.Errorf("GC command not found but '%s'\n", commands[0].Cmd)
	}
}
