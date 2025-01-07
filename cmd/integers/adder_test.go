package integers

import "testing"
import "fmt"

func ExampleAdd() {
    sum := Add(1,5)
    fmt.Println(sum)
    //output 6
}

func TestAdder(t *testing.T) {
    got := Add(2,2)
    want := 4

    if got != want {
        t.Errorf("got: %q, want: %q", got, want)
    }
}
