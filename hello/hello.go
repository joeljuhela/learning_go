package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const finnishHelloPrefix = "Hei, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    prefix := englishHelloPrefix

    switch language {
    case "Finnish":
        prefix = finnishHelloPrefix
    case "Spanish":
        prefix = spanishHelloPrefix
    }
    
    return prefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}
