package src

//栈结构

//使用单链实现栈功能
type Stack struct {
	l *Linked
	size int
}
func NewStack() *Stack {
	return &Stack{
		l:NewLinked(),
		size:0,
	}
}
//push stack
func (s *Stack) Push(data interface{}) {
	s.l.Append(data)
	s.size ++
}
//pop stack
func (s *Stack) Pop() interface{} {
	ret := s.l.RemoveTail()
	if ret != nil {
		s.size --
		return ret
	}
	return nil
}
//Check stack is empty
func (s *Stack) StackEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}
//Get stack length
func (s *Stack) Length() int {
	return s.size
}
//Print linked info
func (s *Stack) Print() string {
	return s.l.Print()
}