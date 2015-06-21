#Prompt [![Build Status](https://travis-ci.org/go-prompt/prompt.svg?branch=master)](https://travis-ci.org/go-prompt/prompt) [![Coverage Status](https://coveralls.io/repos/go-prompt/prompt/badge.svg)](https://coveralls.io/r/go-prompt/prompt) [![GoDoc](https://godoc.org/github.com/go-prompt/prompt?status.svg)](http://godoc.org/github.com/go-prompt/prompt)

This Go package provides a dead simple 'terminal-prompt-like' for your application.


#Install
There is no extra dependencies, just a [Go environment](http://golang.org/doc/install) is needed. Then:
~~~
$ go get gopkg.in/prompt.v1
~~~

#Use
Here you can see an example about how to add a custom command.
~~~ go
package main

import (
    "fmt"
    "gopkg.in/prompt.v1"
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
