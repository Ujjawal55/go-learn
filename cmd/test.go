package main

import "fmt"

func Test(combinedArrays ...[]int) {
    for _, numbers := range combinedArrays {
        fmt.Println(numbers)
    }
}

func main() {
    Test([]int{1,2,3}, []int{3,4,4})
}
