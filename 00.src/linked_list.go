package src

import "fmt"

//单链表

//定义结点
type LinkedNode struct {
	Data interface{} //结点数据
	Next *LinkedNode
}
//实例一个结点
func NewLinkedNode(data interface{}) *LinkedNode {
	return &LinkedNode{
		Data:data,
	}
}
//defined single linked
type Linked struct {
	Head *LinkedNode //head node
	Tail *LinkedNode //tail node
	Size int //record linked length
}
//new linked
func NewLinked() *Linked {
	return &Linked{}
}

//1~10的测试数据
func TestLinkedData() *LinkedNode {
	linked := NewLinked()
	for i := 1; i <= 10; i ++ {
		linked.Append(i)
	}
	return linked.Head
}
//Add a node at the tail
func (l *Linked) Append(data interface{}) {
	node := NewLinkedNode(data)
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
	l.Size ++
}
//Add a node to the header
func (l *Linked) AddHead(data interface{}) {
	node := NewLinkedNode(data)
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		old := l.Head
		l.Head = node
		l.Head.Next = old
	}
	l.Size ++
}
//remove a node
func (l *Linked) Remove(index int) interface{} {
	if l.Head == nil {
		return nil
	}
	if index <= 0 || index > l.Size {
		return nil
	}
	if index == 1 {
		return l.RemoveHead()
	}
	if index == l.Size {
		return l.RemoveTail()
	}
	prev := new(LinkedNode)//创建一个空指针,存放上一次循环的指针,因为单链表没有前趋指针.
	p := l.Head
	for i := 0; i < index-1; i ++{
		prev = p//暂存上一次的指针结点
		p = p.Next//继续下一个指针查找
	}
	prev.Next = p.Next //设置新的结点关系
	l.Size --
	return p.Data
}
//get header data
func (l *Linked) RemoveHead() interface{} {
	if l.Head == nil {
		return nil
	}
	node := l.Head
	l.Head = node.Next
	l.Size --
	if l.Size == 0 {
		l.Head, l.Tail = nil , nil
	}
	return node.Data
}
//删除尾结点
func (l *Linked) RemoveTail() interface{} {
	if l.Head == nil {
		return nil
	}
	//取出尾结点
	node := l.Tail
	//循环找到新的结点做为尾结点
	p := l.Head
	for i := 0; i < l.Size - 2; i ++ {
		p = p.Next
	}
	//找到新的结点,将结点下一个指针设置为nil.做为新尾结点
	p.Next = nil
	l.Tail = p
	l.Size --
	if l.Size == 0 { //当长度为0时,重置头与尾指针
		l.Head, l.Tail = nil, nil
	}
	return node.Data
}

//print linked info
func (l *Linked) Print() string {
	str := ""
	p := l.Head
	for p != nil {
		str += fmt.Sprint(p.Data)
		if p.Next != nil {
			str += "=>"
		}
		p = p.Next
	}
	if str == "" {
		return "linked is null"
	}
	return str
}
//显示结点的详细
func (l *Linked) Display(head *LinkedNode) string {
	str := ""
	p := head
	for p != nil {
		str += fmt.Sprint(p.Data)
		if p.Next != nil {
			str += "=>"
		}
		p = p.Next
	}
	if str == "" {
		return "linked is null"
	}
	return str
}
//反转链表
func (l *Linked) Reserve() *LinkedNode {
	//只有一个结点时,直接返回即可
	if l.Size == 1 {
		return l.Head
	}
	var prev, next *LinkedNode //创建两个空指针,分别是前趋,后续结点
	for l.Head != nil { //从头结点开始循环
		next = l.Head.Next //将头结点的下一个结点,设置为新的头结点.
		l.Head.Next = prev //旧的头结点的下一个结点设置为空或成为前趋结点
		prev = l.Head //将旧的头结点设置为前趋结点
		l.Head = next//将头的下一个结点设置为新的头结点.
	}
	return prev
}

