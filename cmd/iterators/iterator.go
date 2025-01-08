package iterators

const repeatCount = 5
func Repeat(character string, userCount int) string {
    var repeated string
    for i := 1; i <= userCount; i++ {
        repeated = repeated + character
    }

    return repeated
}
