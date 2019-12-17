package linked

import "fmt"

//单链表基本操作
/*
如果链表为空时，代码是否能正常工作？
如果链表只包含一个结点时，代码是否能正常工作？
如果链表只包含两个结点时，代码是否能正常工作？
代码逻辑在处理头结点和尾结点的时候，是否能正常工作？
*/


type SingleNode struct {
	data interface{}
	next *SingleNode
}
//带头链表, 其实head相当于一个哨兵
type SingleLinked struct {
	head   *SingleNode //头部
	length uint
}

func NewSingleNode(data interface{}) *SingleNode {
	return &SingleNode{data:data}
}
func (n *SingleNode) GetData() interface{} {
	return n.data
}
func (n *SingleNode) GetNext() *SingleNode {
	return n.next
}

//实现一个单链表
func NewSingleLinked() *SingleLinked {
	return &SingleLinked{NewSingleNode(0), 0}
}
func (l *SingleLinked) GetLength() uint {
	return l.length
}
//在某个节点后面插入节点
func (l *SingleLinked) InsertAfter(p *SingleNode, v interface{}) bool {
	if p == nil {
		return false
	}
	node := NewSingleNode(v)
	old := p.next
	p.next = node
	node.next = old
	l.length ++
	return true
}
//在某一结点前插入一个结点
func (l *SingleLinked) InsertBefore(p *SingleNode, v interface{}) bool {
	if p == nil {
		return false
	}
	//头部的下一个结点
	cur := l.head.next
	//头部
	pre := l.head
	//循环查找p的位置
	for cur != nil {
		if cur == p {
			break
		}
		//未找到则,继续移动指针
		pre = cur
		cur = cur.next
	}
	//当结点为空时
	if cur == nil {
		return false
	}

	node := NewSingleNode(v)
	pre.next = node
	node.next = cur
	l.length ++
	return true
}
//向头部插入结点
func (l *SingleLinked) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}
//向尾部插入结点
func (l *SingleLinked) InsertToTail(v interface{}) bool {
	cur := l.head //取出头部结点
	for cur.next != nil {//判断是否有下一个结点
		cur = cur.next
	}
	return l.InsertAfter(cur, v)
}
//查找某个位置的结点
func (l *SingleLinked) FindByIndex(index uint) *SingleNode {
	if l.length == 0 || index > l.length {
		return nil
	}
	cur := l.head.next
	var i uint = 0
	for ; i < index; i ++ {
		cur = cur.next
	}
	return cur
}
func (l *SingleLinked) DeleteNode(p *SingleNode) bool {
	if l.length == 0 || p == nil {
		return false
	}
	//查找到p的上一个指针结点
	cur := l.head.next
	pre := l.head
	for cur != nil {
		if cur == p {
			break
		}
		//相当于一个窗口不断的移动
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	pre.next = p.next
	l.length --
	p = nil
	return true
}
//打印指针
func (l *SingleLinked) Print() {
	p := l.head.next
	format := ""
	for p != nil {
		format += fmt.Sprintf("%+v", p.GetData())
		if p.next != nil {
			format += "->"
		}
		p = p.next
	}
	fmt.Println(format)
}
func (l *SingleLinked) Reserve(p *SingleNode) *SingleNode {
	return nil
}