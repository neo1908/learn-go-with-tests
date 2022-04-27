package test_utils

import "testing"

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
