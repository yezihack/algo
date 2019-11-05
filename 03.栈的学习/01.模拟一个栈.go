package main

import "fmt"

type Stack struct {
	Data    [10]int
	Top     int
	MaxSize int
}

func NewStack(size int) *Stack {
	s := Stack{
		MaxSize: size,
		Top:     -1,
	}
	return &s
}
func (s *Stack) isFull() bool {
	if s.Top == s.MaxSize-1 {
		return true
	}
	return false
}
func (s *Stack) isEmpty() bool {
	if s.Top == -1 {
		return true
	}
	return false
}

//提供一个出栈方法
func (s *Stack) Pop() int {
	if s.isEmpty() {
		return -1
	}
	v := s.Data[s.Top]
	s.Top--
	return v
}

//提供一个入栈方法
func (s *Stack) Push(v int) {
	if s.isFull() {
		fmt.Println("stack is full")
		return
	}
	s.Top++
	s.Data[s.Top] = v
}

//打印
func (s *Stack) Print() {
	for i := s.Top; i >= 0; i-- {
		fmt.Print(s.Data[i], ",")
	}
}

func main() {
	s := NewStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push(8)
	s.Push(10)
	s.Print()
}
