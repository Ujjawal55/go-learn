package struct_learning

import "testing"


const pi = 3.141



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
        var got float64 = Area(rectangle)
        var want float64 = 10.0 * 10.0

        if (got != want) {
            t.Errorf("got: %f, want: %f", got, want)
        }
    })

    t.Run("Circle", func(t *testing.T) {
        circle := Circle{10.0}
        got := Area(circle)
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











