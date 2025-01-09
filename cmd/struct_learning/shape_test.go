package struct_learning

import "testing"

func TestPermeter(t *testing.T) {

    // idea here is that we have to define list of test cases therefore follow table driven tests

    const pi = 3.141

    perimetertests := []struct {
        shape Shape
        want float64
    }{
        {Rectangle{10.0, 10.0}, 2.0 * 10.0 * 10.0 },
        {Circle{10.0}, 2.0 * pi * 10.0},
    }

    for _, tt := range perimetertests {
        got := tt.shape.Perimeter()
        want := tt.want

        assertProperty(t, got, want)
    }

    // checkPerimeter := func(t testing.TB, shape Shape, want float64) {
    //     t.Helper()
    //     var got float64 = shape.Perimeter()
    //     assertProperty(t, got, want)
    //
    // }
    //
    // rectangle := Rectangle{10.0, 10.0}
    // want := rectangle.Perimeter()
    // checkPerimeter(t, rectangle, want)
}

func TestArea(t *testing.T) {


    const pi = 3.141
    areaTest := []struct {
        shape Shape
        want float64
    } {
        {Rectangle{10.0, 10.0}, 10.0 * 10.0},
        {Circle{10.0}, pi * 10.0 * 10.0},
    }

    for _, tt := range areaTest {
        got := tt.shape.Area()
        want := tt.want

        assertProperty(t, got, want)
    }




    // checkArea := func(t testing.TB, shape Shape, want float64) {
    //     t.Helper()
    //     got := shape.Area()
    //
    //     assertProperty(t, got, want)
    //
    // }
    //
    // t.Run("Rectangle", func(t *testing.T) {
    //     rectangle := Rectangle{10.0, 10.0}
    //     var want float64 = 10.0 * 10.0
    //     checkArea(t, rectangle, want)
    //
    // })
    //
    // t.Run("Circle", func(t *testing.T) {
    //     const pi = 3.141
    //     circle := Circle{10.0}
    //     want := pi * 10.0 * 10.0
    //     checkArea(t, circle, want)
    // })

}

func assertProperty(t testing.TB, got float64, want float64) {
    t.Helper()
    if (got != want) {
        t.Errorf("got: %f, want: %f", got, want)
    }

}











