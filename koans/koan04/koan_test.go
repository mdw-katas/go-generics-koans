package koan04

import "testing"

func Test(t *testing.T) {
	t.Log("Focus: What about custom slice types?")
	t.Log("https://pkg.go.dev/slices")

	if !Contains([]int{1, 2, 3}, 2) {
		t.Fatal("didn't find expected value")
	}
	if !Contains([]string{"a", "b", "c"}, "b") {
		t.Fatal("didn't find expected value")
	}
	if !Contains([]rune("asdf"), 'f') {
		t.Fatal("didn't find expected value")
	}
	if !Contains(CustomSlice{true, false}, true) {
		t.Fatal("didn't find expected value")
	}
}

type CustomSlice []bool

// DO NOT MODIFY ANY CODE ABOVE THIS LINE ///////////
/////////////////////////////////////////////////////
// Mess with the code below to make all tests pass //

func Contains[V comparable, S []V](slice S, value V) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Scroll down for a hint...
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// Hint: all that is needed is the addition of a single character...
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// HINT: all you need to add is a ~ character.
