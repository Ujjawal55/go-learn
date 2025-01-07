package main

import "testing"

func TestHello(t *testing.T) {

    t.Run("Greeting in Spaninsh", func (t *testing.T) {
        var got string = Hello("Kash", "French")
        var want string = "Bonjour, Kash"

        assertCorrectMessage(t, got, want)
    })
}

// code refactoring for (don't repeat yourself principle)
func assertCorrectMessage(t testing.TB, got string, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got: %q, want: %q",got, want)
    }
}

