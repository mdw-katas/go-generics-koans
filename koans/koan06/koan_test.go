package koan06

import (
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	t.Log("https://pkg.go.dev/container/ring")

	// int values:
	intRing := New[int](5)
	for i := 0; i < 5; i++ {
		intRing.Value = i
		intRing = intRing.Next()
	}
	intRing.Do(func(v int) { t.Log(v) })

	// string values:
	stringRing := New[string](5)
	for i := 0; i < 5; i++ {
		stringRing.Value = fmt.Sprint(i)
		stringRing = stringRing.Next()
	}
	stringRing.Do(func(v string) { t.Log(v) })
}

// DO NOT MODIFY ANY CODE ABOVE THIS LINE ///////////
/////////////////////////////////////////////////////
// Mess with the code below to make all tests pass //

type Ring struct {
	next, prev *Ring
	Value      any
}

func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}
func (this *Ring) init() *Ring {
	this.next = this
	this.prev = this
	return this
}
func (this *Ring) Next() *Ring {
	if this.next == nil {
		return this.init()
	}
	return this.next
}
func (this *Ring) Prev() *Ring {
	if this.next == nil {
		return this.init()
	}
	return this.prev
}
func (this *Ring) Move(n int) *Ring {
	if this.next == nil {
		return this.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			this = this.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			this = this.next
		}
	}
	return this
}
func (this *Ring) Link(s *Ring) *Ring {
	n := this.Next()
	if s != nil {
		p := s.Prev()
		this.next = s
		s.prev = this
		n.prev = p
		p.next = n
	}
	return n
}
func (this *Ring) Unlink(n int) *Ring {
	if n <= 0 {
		return nil
	}
	return this.Link(this.Move(n + 1))
}
func (this *Ring) Len() int {
	n := 0
	if this != nil {
		n = 1
		for p := this.Next(); p != this; p = p.next {
			n++
		}
	}
	return n
}
func (this *Ring) Do(f func(any)) {
	if this != nil {
		f(this.Value)
		for p := this.Next(); p != this; p = p.next {
			f(p.Value)
		}
	}
}
