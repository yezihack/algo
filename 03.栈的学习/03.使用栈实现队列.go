package main

import (
	"fmt"
	 "github.com/yezihack/algo/00.src"
)

//栈实现队列效果
//栈是先进后出
//队列是先进先出.
//解题思路: 使用两个栈实现,一个叫push栈,负责入栈操作,一个叫peek栈,负责出栈操作
//当出栈操作时,将push栈的压入peek栈中.从而实现队列效果

//定义一个栈队列结构
type StackQueue struct {
	PushStack *src.Stack//push栈
	PeekStack *src.Stack//peek栈
}
func NewStackQueue() *StackQueue {
	return &StackQueue{
		PushStack:src.NewStack(),
		PeekStack:src.NewStack(),
	}
}
//入队列操作
func (q *StackQueue) Push(data interface{}) {
	q.PushStack.Push(data)
}
//出队列操作
func (q *StackQueue) Pull() interface{} {
	//先判断peekStack是否为空,如果空则把pushStack栈的元素搬到peekStack栈中
	if !q.PeekStack.StackEmpty() {
		return q.PeekStack.Pop()
	}
	size := q.PushStack.Length()
	//for i := 0; i < q.PushStack.Length(); i ++ {} //代码1处
	for i := 0; i < size; i ++ {
		q.PeekStack.Push(q.PushStack.Pop())
	}
	return q.PeekStack.Pop()
}
func main() {
	q := NewStackQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pull())
	fmt.Println(q.Pull())
	fmt.Println(q.Pull())
}