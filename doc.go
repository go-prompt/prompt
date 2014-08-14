/*
Package prompt provides a 'terminal-prompt-like' for your application.

Usage

See an example about how to add a custom command:

	package main

	import (
	    "fmt"
	    "github.com/viliamjr/prompt"
	)

	func main() {

	    prompt.AddCommand(
	        "ping",
	        "my custom command",
	        func() {
	            fmt.Println("pong!")
	        },
	    )

	    <-prompt.Start()
	    // if you don't need to lock, just remove the '<-' operator

	    fmt.Println("Bye from main!")
	}

This will show a prompt, where you can try your custom command 'ping'.
Or just type 'quit' to close it. For a list of commands type 'help'.
*/
package prompt
