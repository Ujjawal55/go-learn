package struct_learning

import "testing"

func TestPermeter(t *testing.T) {

    checkPerimeter := func(t testing.TB, shape Shape, want float64) {
        t.Helper()
        var got float64 = shape.Perimeter()
        assertProperty(t, got, want)

    }

    rectangle := Rectangle{10.0, 10.0}
    want := rectangle.Perimeter()
    checkPerimeter(t, rectangle, want)
}

func TestArea(t *testing.T) {

    checkArea := func(t testing.TB, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()

        assertProperty(t, got, want)

    }

    t.Run("Rectangle", func(t *testing.T) {
        rectangle := Rectangle{10.0, 10.0}
        var want float64 = 10.0 * 10.0
        checkArea(t, rectangle, want)

    })

    t.Run("Circle", func(t *testing.T) {
        const pi = 3.141
        circle := Circle{10.0}
        want := pi * 10.0 * 10.0
        checkArea(t, circle, want)
    })

}

func assertProperty(t testing.TB, got float64, want float64) {
    t.Helper()
    if (got != want) {
        t.Errorf("got: %f, want: %f", got, want)
    }

}











