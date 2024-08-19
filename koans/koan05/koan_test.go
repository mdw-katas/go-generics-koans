package koan05

import "testing"

func Test(t *testing.T) {
	t.Log("Focus: Map functions")
	t.Log("https://pkg.go.dev/maps")
}
func TestSimpleMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	Clear(m)
	if len(m) > 0 {
		t.Fatal("should be empty")
	}
}
func TestCustomMap(t *testing.T) {
	m := MyMap{1: "one", 2: "two"}
	Clear(m)
	if len(m) > 0 {
		t.Fatal("should be empty")
	}
}

type MyMap map[int]string

// DO NOT MODIFY ANY CODE ABOVE THIS LINE ///////////
/////////////////////////////////////////////////////
// Mess with the code below to make all tests pass //

func Clear(m map[int]string) {
	for k := range m {
		delete(m, k)
	}
}
