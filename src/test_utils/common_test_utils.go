package test_utils

import (
	"reflect"
	"testing"
)

func ValidateTestState(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

func ValidateIntTestState(got int, want int, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func ValidateSliceTestState(got []int, want []int, t testing.TB) {
	t.Helper()

	// Cannot use equality ops on slices so need to use reflect.DeepEqual
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
