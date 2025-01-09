package struct_learning

import "testing"

func TestPermeter(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := Perimeter(rectangle)
    want := 2.0 * 10.0 * 10.0

    if got != want {
        t.Errorf("got: %v, want: %f", got, want)
    }
}

func TestArea(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    var got float64 = Area(rectangle)
    var want float64 = 10.0 * 10.0

    if (got != want) {
        t.Errorf("got: %f, want: %f", got, want)
    }

}
