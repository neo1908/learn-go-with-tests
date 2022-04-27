package src

import (
	"learn-go-with-tests/src/test_utils"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	want := 4

	test_utils.ValidateIntTestState(sum, want, t)
}
