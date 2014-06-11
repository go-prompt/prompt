package prompt

import (
	"bufio"
	"fmt"
	"os"
)

var commands []Command

func start(control chan bool) {

	scanner := bufio.NewScanner(os.Stdin)

	var stoped bool

	for !stoped {

		fmt.Print("> ")
		if scanner.Scan() {
			command := Parse(scanner.Text())
			stoped = command.Run()
		} else { // CONTROL-C
			stoped = true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	control <- true
}

func Start() chan bool {

	control := make(chan bool)
	go start(control)
	return control
}

func Add(c Command) {

	commands = append(commands, c)
}

func AddCommand(name string, desc string, action func()) {

	commands = append(commands, Command{name, desc, action})
}
