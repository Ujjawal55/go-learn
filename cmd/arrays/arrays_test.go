package arrays


import "testing"

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

func assertCorrectionAnswer(t testing.TB, got int, want int) {
    t.Helper()
    if got != want {
        t.Errorf("got: %v, want: %v", got, want)
    }
}
