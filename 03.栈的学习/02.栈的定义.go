package main

import (
	"code.qschou.com/golang/errors"
	"fmt"
)

//栈是一种后进先出LILO的线性数据类型
type Object int
type LiloStacker interface {
	IsEmpty() bool          //判断栈是否为空
	IsFull() bool           //判断栈是否满了
	StackLength() int       //获取一个栈的长度
	GetTop() Object         //获取栈顶元素
	ClearStack()            //清空栈
	Push(data Object) error //压栈
	Pop() Object            //出栈
}
type LILOStack struct {
	base int
	top  int
	data [3]Object
}

func NewLILOStack() *LILOStack {
	return &LILOStack{}
}
func (s *LILOStack) Push(data Object) error {
	//is stack full ?
	if s.IsFull() {
		return errors.New("stack is full")
	}
	s.data[s.top] = data
	s.top++
	return nil
}
func (s *LILOStack) GetTop() Object {
	if s.IsEmpty() {
		return -1
	}
	return s.data[len(s.data)-1]
}
func (s *LILOStack) Pop() Object {
	if s.IsEmpty() {
		return -1
	}
	index := s.top - 1
	s.top--
	return s.data[index]
}
func (s *LILOStack) ClearStack() {
	s.top = s.base
}
func (s *LILOStack) StackLength() int {
	return s.top - s.base
}

func (s *LILOStack) IsFull() bool {
	if (s.top - s.base) == len(s.data) {
		return true
	}
	return false
}
func (s *LILOStack) IsEmpty() bool {
	if s.top == s.base {
		return true
	}
	return false
}

func main() {
	var s LiloStacker
	s = NewLILOStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s.Pop())
	fmt.Println(s.StackLength())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
