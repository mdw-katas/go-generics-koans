package koan03

import "testing"

func Test(t *testing.T) {
	t.Log("Focus: generics make it possible to operate on slices of any item type.")
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
}

// DO NOT MODIFY ANY CODE ABOVE THIS LINE ///////////
/////////////////////////////////////////////////////
// Mess with the code below to make all tests pass //

func Contains(slice []any, value any) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
