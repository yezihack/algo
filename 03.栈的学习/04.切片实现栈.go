package main

import "fmt"

type StackEntry struct {
	items []interface{}
	length int
	cap int
}
func NewStackEntry(cap int) *StackEntry {
	return &StackEntry{
		items:make([]interface{}, cap, cap),
		cap:cap,
	}
}
func (s *StackEntry) Push(v interface{}) bool {
	if s.length == s.cap {
		return false
	}
	s.items[s.length] = v
	s.length ++
	return true
}
func (s *StackEntry) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	v := s.items[s.length-1]
	s.items[s.length-1] = nil
	s.length --
	return v
}
func main() {
	s := NewStackEntry(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Push("a"))
	fmt.Println(s.Push("b"))
	fmt.Println(s.Push("c"))
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
