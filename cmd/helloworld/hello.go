package main

import (
	"fmt"
)

const (
    spanish = "Spanish"
    french = "French"

    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Halo, "
    frenchHelloPrefix = "Bonjour, "
)

func Hello(name, language string) string {
    if name == "" {
        name = "World"
    }

    return grettingPrefix(language) + name
}

func grettingPrefix(language string) (prefix string) {

    switch language {
    case spanish:
        prefix = spanishHelloPrefix
    case french:
        prefix = frenchHelloPrefix
    default:
        prefix = englishHelloPrefix
    }

    return
}


func main() {
    fmt.Println(Hello("", "Spanish"))
}




















