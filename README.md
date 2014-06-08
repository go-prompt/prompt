#What?
Pimp your Go app with a dead simple prompt feature.

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

    my_cmd := got.Command{
        "ping",
        "just an example",
        func() {
            fmt.Println("pong!")
        },
    }

    got.Add(my_cmd)

    <-got.Start()
    // if you don't need to lock, just remove the '<-' operator

    fmt.Println("Bye from main!")
}
~~~

This will show a prompt, where you can try your custom command 'ping'. Or just type 'quit' to close it.

#TODO
- [ ] help/list commands
- [ ] some documentation in golang format
- [ ] builtin commands (memory,threads,network?)
