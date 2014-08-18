#Prompt [![Build Status](https://travis-ci.org/viliamjr/prompt.svg?branch=master)](https://travis-ci.org/viliamjr/prompt) [![GoDoc](https://godoc.org/github.com/viliamjr/prompt?status.svg)](http://godoc.org/github.com/viliamjr/prompt)

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

This example will show a _terminal-prompt-shell-like_, where you can try your custom command '**ping**'. Or just type '**quit**' to close it. For a list of commands type '**help**'. You can see this in action **[here](https://asciinema.org/a/11556)**.

#TODO
- [ ] a wrapper for _exec.Command_ (from _os_ standard library).
