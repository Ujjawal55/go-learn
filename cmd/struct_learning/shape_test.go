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
    t.Run("Rectangle", func(t *testing.T) {
        rectangle := Rectangle{10.0, 10.0}
        var got float64 = rectangle.Area()
        var want float64 = 10.0 * 10.0

        if (got != want) {
            t.Errorf("got: %f, want: %f", got, want)
        }
    })

    t.Run("Circle", func(t *testing.T) {
        const pi = 3.141
        circle := Circle{10.0}
        got := circle.Area()
        want := pi * 10.0 * 10.0

        assertArea(t, got, want)
    })

}


func assertArea(t testing.TB, got float64, want float64) {
    t.Helper()
    if (got != want) {
        t.Errorf("got: %f, want: %f", got, want)
    }

}











