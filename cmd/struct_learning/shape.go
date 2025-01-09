package struct_learning

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    width float64
    height float64
}

func (r Rectangle) Area() float64 {
    return r.height * r.width
}


func (r Rectangle) Perimeter() float64 {
    return 2.0 * r.height * r.width
}

type Circle struct {
    radius float64
}


func (c Circle) Area() float64 {
    const pi = 3.141
    return pi * c.radius * c.radius
}


func (c Circle) Perimeter() float64 {
    const pi = 3.141
    return 2.0 * pi * c.radius
}



// cannot define multiple Area function

// func Area(rectangle Rectangle) float64 {
//     return rectangle.width * rectangle.height
// }
//
// func Area(circle Circle) float64 {
//     return pi * circle.radius * circle.radius
// }
