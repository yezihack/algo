package src

import (
	"fmt"
	"strconv"
)

type Element int

//创建一个单链表

//创建一个结点
type SingleNode struct {
	Next *SingleNode
	Data Element
}

//新建一个结点
func NewSingleNode(data Element) *SingleNode {
	return &SingleNode{
		Data: data,
	}
}

type SingleLinker interface {
	Insert(index int, node *SingleNode) //任意位置插入一个结点
	AddHead(node *SingleNode)           //添加一个头结点
	Append(node *SingleNode)            //追加一个结点
	IsEmpty() bool                      //判断链表是否为空
}

//定义一个单链表结构
type SingleLink struct {
	Head *SingleNode //头结点
	Len  int         //长度
}

//实例一个空单链表
func NewSingleLink() *SingleLink {
	return &SingleLink{}
}

//判断链表是否为空
func (s *SingleLink) IsEmpty() bool {
	if s.Head == nil {
		return true
	}
	return false
}

//追加一个结点,相当于向尾部添加一个结点
func (s *SingleLink) Append(node *SingleNode) {
	if s.IsEmpty() {
		s.Head = node
	} else {
		cur := s.Head         //获取头结点
		for cur.Next != nil { //判断链表是否为空,为空则是尾结点,不为空直到遍历到下一个指针为空,则找到尾结点
			cur = cur.Next
		}
		cur.Next = node //找到尾结点,则将新结点设置到旧的尾结点
	}
	s.Len++
}

//添加头部结点
func (s *SingleLink) AddHead(node *SingleNode) {
	if s.IsEmpty() {
		s.Head = node
	} else {
		cur := s.Head
		s.Head = node
		s.Head.Next = cur
	}
	s.Len++
}

//插入链表的位置, 需要判断是否是头结点,是否是尾结点, 是否是中间位置
func (s *SingleLink) Insert(index int, node *SingleNode) {
	if index <= 0 { //头部位置
		s.AddHead(node)
	} else if index > s.Len { //尾部位置
		s.Append(node)
	} else { //中间位置
		cur := s.Head
		for i := 0; i < (index - 1); {
			cur = cur.Next
			i++
		}
		next := cur.Next
		cur.Next = node
		node.Next = next
	}
}
func (s *SingleLink) Print() {
	p := s.Head
	str := "Link-Length:" + strconv.Itoa(s.Len)
	str += "\n"
	for p != nil {
		str += fmt.Sprint(p.Data)
		p = p.Next
		if p != nil {
			str += "=>"
		}
	}
	fmt.Println(str)
}
