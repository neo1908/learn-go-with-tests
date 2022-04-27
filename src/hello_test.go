package src

import (
	"learn-go-with-tests/src/test_utils"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("Say hello to everyone", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		test_utils.ValidateTestState(got, want, t)
	})
}

func TestHelloName(t *testing.T) {

	t.Run("Say hello to Ben", func(t *testing.T) {
		got := Hello("Ben")
		want := "Hello, Ben!"

		test_utils.ValidateTestState(got, want, t)
	})
}
