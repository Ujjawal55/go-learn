package main

import "fmt"

const englishHelloPrefix = "Hello, "
func Hello(name, language string) string {
    if name == "" {
        name = "World"
    }

    if language == "Spanish" {
        return "Halo, " + name
    }

    return englishHelloPrefix  + name
}

func main() {
    fmt.Println(Hello("", "Spanish"))
}

