package struct_learning

type Rectangle struct {
    width float64
    height float64
}

func Perimeter(rectangle Rectangle) float64 {
    return 2.0 * rectangle.width * rectangle.height
}

func Area(rectangle Rectangle) float64 {
    return rectangle.width * rectangle.height
}
