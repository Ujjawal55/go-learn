package main

import "testing"


func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, World"

    if got != want {
        t.Errorf("got =  %v, want = %v", got, want)
    }

}

func TestGreeting(t *testing.T) {
    got := Greeting("kash")
    want := "Hello, kash"

    if got != want {
        t.Errorf("got: %q, want: %q", got, want)
    }

}
