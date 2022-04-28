package src

import (
	"learn-go-with-tests/src/test_utils"
	"testing"
)

// Interface resolution is implicit in Go. Any type that defines a method called Area that returns a float64
// will satisfy the interface. Circle does but string doesn't
type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	test_utils.ValidateFloatTestState(got, want, t)
}

func TestArea(t *testing.T) {

	rectangle := Rectangle{12.0, 6.0}
	want := 72.0

	CheckArea(t, rectangle, want)
}

func TestCircleArea(t *testing.T) {

	circle := Circle{10}
	want := 314.1592653589793

	CheckArea(t, circle, want)
}

func CheckArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()

	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}
