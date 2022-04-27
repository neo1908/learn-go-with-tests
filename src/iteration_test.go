package src

import (
	"fmt"
	"learn-go-with-tests/src/test_utils"
	"testing"
)

func TestRepeat5(t *testing.T) {

	t.Run("Repeat a letter 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		test_utils.ValidateTestState(got, want, t)
	})
}

func TestRepeat300(t *testing.T) {

	t.Run("Repeat a letter 300 times", func(t *testing.T) {
		got := Repeat("z", 300)
		want := "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"

		test_utils.ValidateTestState(got, want, t)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("E", 5)
	}
}

func ExampleRepeat() {

	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
