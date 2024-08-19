package koan02

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Focus: Constructor function and method for a generic type.")
	t.Log("https://tip.golang.org/ref/spec#Type_constraints")

	t.Log(NewPoint(int(1), int(2)))
	t.Log(NewPoint(int8(1), int8(2)))
	t.Log(NewPoint(int16(1), int16(2)))
	t.Log(NewPoint(int32(1), int32(2)))
	t.Log(NewPoint(int64(1), int64(2)))
	t.Log(NewPoint(uint(1), uint(2)))
	t.Log(NewPoint(uint8(1), uint8(2)))
	t.Log(NewPoint(uint16(1), uint16(2)))
	t.Log(NewPoint(uint32(1), uint32(2)))
	t.Log(NewPoint(uint64(1), uint64(2)))
	t.Log(NewPoint(uintptr(1), uintptr(2)))
	t.Log(NewPoint(byte(1), byte(2))) // alias!
	t.Log(NewPoint(rune(1), rune(2))) // alias!
	t.Log(NewPoint(float32(1), float32(2)))
	t.Log(NewPoint(float64(1), float64(2)))
	t.Log(NewPoint(complex64(1), complex64(2)))
	t.Log(NewPoint(complex128(1), complex128(2)))
}

type (
	Number interface {
		Integer | Float | Complex
	}
	Integer interface {
		Signed | Unsigned
	}
	Signed interface {
		int | int8 | int16 | int32 | int64
	}
	Unsigned interface {
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
	}
	Float interface {
		float32 | float64
	}
	Complex interface {
		complex64 | complex128
	}
)

// DO NOT MODIFY ANY CODE ABOVE THIS LINE ///////////
/////////////////////////////////////////////////////
// Mess with the code below to make all tests pass //

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

type Point struct {
	X, Y int
}

func (this Point) String() string {
	return fmt.Sprintf("(%v, %v)", this.X, this.Y)
}
