package src

import (
	"learn-go-with-tests/src/test_utils"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 nums", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		test_utils.ValidateIntTestState(got, want, t)
	})
}

func TestSumWithSlice(t *testing.T) {

	t.Run("Collection of any size -- slize", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		test_utils.ValidateIntTestState(got, want, t)
	})
}

func TestSumAll(t *testing.T) {

	t.Run("Test summing several slices returning a slice of sums", func(t *testing.T) {

		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		test_utils.ValidateSliceTestState(got, want, t)
	})
}

func TestSumAllTails(t *testing.T) {

	t.Run("Test a sum of all tails (All elements except the first)", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		test_utils.ValidateSliceTestState(got, want, t)
	})
}

func TestSumAllTailsWithEmptySlice(t *testing.T) {

	t.Run("Test a sum of all tails with empty slice", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		test_utils.ValidateSliceTestState(got, want, t)
	})
}
