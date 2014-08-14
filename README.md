#Prompt [![Build Status](https://travis-ci.org/viliamjr/prompt.svg?branch=master)](https://travis-ci.org/viliamjr/prompt)

This Go package provides a dead simple 'terminal-prompt-like' for your application.

#Install
~~~
$ go get github.com/viliamjr/prompt
~~~

#Use
Here you can see an example about how to add a custom command.
~~~ go
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
~~~

This will show a prompt, where you can try your custom command 'ping'. Or just type 'quit' to close it. For a list of commands type 'help'.

#TODO
- [ ] a wrapper for _exec.Command_ (from _os_ standard library).
