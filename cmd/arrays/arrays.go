package arrays


func SumAllTails(arraysOfNumbers ...[]int) []int{
    var lengthOfNumbers int = len(arraysOfNumbers)
    sums := make([]int, lengthOfNumbers)

    for idx, numbers := range arraysOfNumbers {
        sums[idx] += SumTails(numbers)
    }

    return sums
}


func SumTails(numbers []int) int {
    lengthOfnumbers := len(numbers)
    var sum int = 0
    for i := 1; i < lengthOfnumbers; i++ {
        sum += numbers[i]
    }

    return sum
}

func SumAll(numbersToSum ...[]int) ([]int) {

    var sums []int

    for _, number := range numbersToSum {
        sums = append(sums, Sum(number))
    }
    return sums
}

func Sum(numbers []int) int {
    var sum int = 0

    for _, num := range numbers {
        sum += num
    }
    return sum
}
