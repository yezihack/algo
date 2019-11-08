package main

//defined node struct
type StackTreeNode struct {
	data *BiNode        //save data
	prev *StackTreeNode //front node
	next *StackTreeNode //rear node
}

//init node
func NewStackNode(data *BiNode) *StackTreeNode {
	return &StackTreeNode{
		data: data,
	}
}

//defined stack struct
type Stack struct {
	head *StackTreeNode //head node
	tail *StackTreeNode //tail node
	len  int            //record length
}

//init stack
func InitStack() *Stack {
	return &Stack{}
}

//push stack
func (s *Stack) Push(data *BiNode) {
	newNode := NewStackNode(data)
	if s.len == 0 {
		s.head = newNode
		s.tail = s.head
	} else {
		node := s.head
		s.head = newNode
		s.head.next = node
	}
	s.len++
}

//pop stack
func (s *Stack) Pop() (data *BiNode) {
	if s.len == 0 {
		return
	} else {
		node := s.head
		s.head = node.next
		data = node.data
	}
	s.len--
	return
}

//check stack is empty
func (s *Stack) StackEmpty() bool {
	if s.len == 0 {
		return true
	}
	return false
}
