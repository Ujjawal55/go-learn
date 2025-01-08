package iterators

import "testing"
import "fmt"

func ExampleRepeat() {
    var result string = Repeat("a", 4)
    fmt.Printf("%q", result)
    // output: "aaaa"
}


func TestRepeat(t *testing.T) {
    got := Repeat("a", 5)
    want := "aaaaa"

    if got != want {
        t.Errorf("got: %q, want: %q", got, want)
    }
}
