package src

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	So(t, 1, s.Pop())
	So(t, 2, s.Pop())
	So(t, 3, s.Pop())
	So(t, nil, s.Pop())
	s.Push(8)
	So(t, 8, s.Pop())
}
func TestStack_Length(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	So(t, 3, s.Length())
	fmt.Println("pop", s.Pop())
	So(t, 2, s.Length())
	fmt.Println("pop", s.Pop())
	So(t, 1, s.Length())
	fmt.Println("pop", s.Pop())
	So(t, 0, s.Length())
	fmt.Println("pop", s.Pop())
	So(t, 0, s.Length())
}