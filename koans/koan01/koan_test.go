package koan01

import "testing"

func Test(t *testing.T) {
	t.Log("Focus: Type declarations now accept type parameters.")
	t.Log("https://tip.golang.org/ref/spec#Type_parameter_declarations")

	t.Log(Point[float64]{X: 1, Y: 2})
	t.Log(Point[float32]{X: 1, Y: 2})
	t.Log(Point[uintptr]{X: 1, Y: 2})
}

// DO NOT MODIFY ANY CODE ABOVE THIS LINE ///////////
/////////////////////////////////////////////////////
// Mess with the code below to make all tests pass //

type Point struct {
	X, Y float64
}
