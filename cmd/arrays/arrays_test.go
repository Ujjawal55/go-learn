package arrays

import "testing"
import "reflect"


func TestSumAllTails(t *testing.T) {

    got := SumAllTails([]int{0,2}, []int{3,4})
    want := []int{2,4}

    assertTwoEqualArrays(t, got, want)
}


func TestSumALL(t *testing.T) {
    num1 := []int{1,2}
    num2 := []int{3,4}

    got := SumAll(num1, num2)
    want := []int{3, 7}
    assertTwoEqualArrays(t, got, want)
}

func TestSum(t *testing.T) {
    // Run("fixed capacity array", func (t *testing.T) {
    //     numbers := [5]int{1,2,3,4,5}
    //
    //     got := Sum(numbers)
    //     want := 15
    //
    //     assertCorrectionAnswer(t, got, want)
    // })

    t.Run("variable size array(slice)", func (t *testing.T) {
        var numbers []int = []int{1,2,3,4,5}

        var got int = Sum(numbers)
        var want int = 15

        assertCorrectionAnswer(t, got, want)
    })
}

func assertTwoEqualArrays(t testing.TB, got []int, want []int) {
    t.Helper()
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got: %v, want: %v", got, want)
    }
}

func assertCorrectionAnswer(t testing.TB, got int, want int) {
    t.Helper()
    if got != want {
        t.Errorf("got: %v, want: %v", got, want)
    }
}














