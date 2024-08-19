package koan00

import "testing"

func Test(t *testing.T) {
	t.Log("Focus: Function definitions now accept type parameters.")
	t.Log("https://tip.golang.org/ref/spec#Type_parameter_declarations")
	t.Log("https://tip.golang.org/ref/spec#Function_declarations")
}

func TestInt(t *testing.T) {
	a, b := int(1), int(2)
	m := Min(a, b)
	if m != a {
		t.Fatalf("expected %d, got %d", a, m)
	}
}
func TestUInt(t *testing.T) {
	a, b := uint(1), uint(2)
	m := Min(b, a)
	if m != a {
		t.Fatalf("expected %d, got %d", a, m)
	}
}

// DO NOT MODIFY ANY CODE ABOVE THIS LINE ///////////
/////////////////////////////////////////////////////
// Mess with the code below to make all tests pass //

// Min returns the smallest of the provided arguments.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
