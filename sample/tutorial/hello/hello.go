package main

import (
    "fmt"
    "log"
    "rsc.io/quote"
    "example.com/greetings"
)

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())

    log.SetPrefix("greetings: ")
    log.SetFlags(0)
    message, err := greetings.Hello("Gladys")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)

    names := []string{"Gladys", "Samantha", "Darrin"}
    messages, _ := greetings.Hellos(names)
    fmt.Println(messages)
}

