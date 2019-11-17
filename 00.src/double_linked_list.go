package src

import "fmt"

//双向链表
type DLinkedNode struct {
	Data interface{} //结点数据
	Prev *DLinkedNode//前趋结点
	Next *DLinkedNode//后续结点
}
func NewDLinkedNode(data interface{}) *DLinkedNode {
	return &DLinkedNode{
		Data:data,
	}
}
type DLinkedList struct {
	Head *DLinkedNode //头结点
	Tail *DLinkedNode //尾结点
	size int //链表长度
}
func NewDLinkedList() *DLinkedList {
	return &DLinkedList{}
}
//追加结点
func (l *DLinkedList) Append(data interface{}) {
	node := NewDLinkedNode(data)
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		old := l.Tail
		l.Tail = node
		l.Tail.Prev = old
		old.Next = l.Tail
	}
	l.size ++
}
//头部添加
func (l *DLinkedList) AddHead(data interface{}) {
	node := NewDLinkedNode(data)
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		old := l.Head
		l.Head = node
		l.Head.Next = old
		old.Prev = l.Head
	}
	l.size ++
}
//删除头结点
func (l *DLinkedList) RemoveHead() interface{} {
	if l.Head == nil {
		return nil
	}
	node := l.Head
	if l.size == 1 {
		l.Head, l.Tail = nil, nil
		return node.Data
	}
	l.Head = node.Next
	l.Head.Prev = nil
	l.size --
	if l.size == 0 {
		l.Head, l.Tail = nil, nil
	}
	return node.Data
}
//删除尾结点
func (l *DLinkedList) RemoveTail() interface{} {
	if l.Head == nil {
		return nil
	}

	node := l.Tail
	if l.size == 1 {
		l.Head, l.Tail = nil, nil
		return node.Data
	}
	l.Tail = node.Prev
	l.Tail.Next = nil
	l.size --
	if l.size == 0 {
		l.Head, l.Tail = nil, nil
	}
	return node.Data
}
//删除任意结点
func (l *DLinkedList) Remove(index int) interface{} {
	if l.Head == nil {
		return nil
	}
	//判断是否越界
	if index <= 0 || index > l.size {
		return nil
	}
	//判断是否是头结点
	if index == 1 {
		return l.RemoveHead()
	}
	//判断是否是尾结点
	if index == l.size {
		return l.RemoveTail()
	}
	p := l.Head
	for i := 0; i < index;i ++ {
		p = p.Next
	}
	p.Prev.Next = p.Next
	p.Next.Prev = p.Prev
	l.size --
	if l.size == 0 {
		l.Head, l.Tail = nil, nil
	}
	return p.Data
}
//获取链表长度
func (l *DLinkedList) Length() int {
	return l.size
}
//打印链表
func (l *DLinkedList) Print() string {
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
//显示链表详细
func (l *DLinkedList) Display(head *DLinkedNode) string {
	str := ""
	p := head
	for p != nil {
		str += fmt.Sprint(p.Data)
		if p.Next != nil {
			str += "=>"
		}
		p = p.Next
	}
	return str
}
//反转链表
func (l *DLinkedList) Reverse() *DLinkedNode {
	var prev, next *DLinkedNode
	for l.Head != nil {
		next = l.Head.Next
		l.Head.Next = prev
		prev = l.Head
		l.Head = next
	}
	return prev
}