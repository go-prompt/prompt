package prompt

import (
	"fmt"
	"runtime"
)

type Command struct {
	Cmd    string
	Desc   string
	Action func()
}

func (c *Command) Run() bool {

	if c.Cmd == "quit" {
		return true
	}

	builtinCommand(c)

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

func builtinCommand(c *Command) {
	if c.Cmd == "" {
		c.Action = func() {}
	} else if c.Cmd == "runtime" || c.Cmd == "r" {
		c.Action = func() {
			var s runtime.MemStats
			runtime.ReadMemStats(&s)

			fmt.Println("###")
			fmt.Println("# Logical CPUs:", runtime.NumCPU())
			fmt.Println("# Goroutines  :", runtime.NumGoroutine())
			fmt.Println("# Memory")
			fmt.Println("# |- Allocated (bytes)")
			fmt.Println("# | |- General:", s.Alloc)
			fmt.Println("# | |- Heap   :", s.HeapAlloc)
			fmt.Println("# |- Number of")
			fmt.Println("#   |- Mallocs:", s.Mallocs)
			fmt.Println("#   |- Frees  :", s.Frees)
			fmt.Println("###")
		}
	}
}
