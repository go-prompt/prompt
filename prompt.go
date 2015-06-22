package prompt

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Start initialize the prompt in a new goroutine.
// Returns a boolean channel to control/synchronize your app flow (if appropriate).
func Start() chan bool {

	control := make(chan bool)
	go start(control)
	return control
}

func start(control chan bool) {

	builtinCommands()

	scanner := bufio.NewScanner(os.Stdin)

	var stoped bool

	for !stoped {

		fmt.Print("> ")
		if scanner.Scan() {
			args := strings.Split(scanner.Text(), " ")
			command := findCommand(args[0])
			stoped = command.run(args[1:])
		} else { // CONTROL-C
			stoped = true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	control <- true
}

func builtinCommands() {

	AddCommand("help", "well, I guess you already know.",
		func(args []string) {
			for _, c := range commands {
				fmt.Printf("%20s - %s\n", c.Cmd, c.Desc)
			}
		})

	AddCommand("runtime", "display runtime information.",
		func(args []string) {
			var s runtime.MemStats
			runtime.ReadMemStats(&s)

			fmt.Println()
			fmt.Println(" Logical CPUs:", runtime.NumCPU())
			fmt.Println(" Goroutines  :", runtime.NumGoroutine())
			fmt.Println(" Memory")
			fmt.Println("  * Allocated (bytes)")
			fmt.Println("   - general:", s.Alloc)
			fmt.Println("   - heap   :", s.HeapAlloc)
			fmt.Println("  * Number of")
			fmt.Println("   - mallocs:", s.Mallocs)
			fmt.Println("   - frees  :", s.Frees)
			fmt.Println()
		})

	AddCommand("gc", "call garbage collector.",
		func(args []string) {
			runtime.GC()
			fmt.Println("GC called.")
		})

	AddCommand("quit", "close prompt.", func(args []string) {})
}
