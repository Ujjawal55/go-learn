package struct_learning



type Rectangle struct {
    width float64
    height float64
}

func (r Rectangle) Area() float64 {
    return r.height * r.width
}

type Circle struct {
    radius float64
}


func (c Circle) Area() float64 {
    const pi = 3.141
    return pi * c.radius * c.radius
}


func Perimeter(rectangle Rectangle) float64 {
    return 2.0 * rectangle.width * rectangle.height
}

// cannot define multiple Area function

// func Area(rectangle Rectangle) float64 {
//     return rectangle.width * rectangle.height
// }
//
// func Area(circle Circle) float64 {
//     return pi * circle.radius * circle.radius
// }
