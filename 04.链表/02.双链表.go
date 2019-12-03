package linked

import "fmt"

type Element interface{}

type DoubleNode struct {
	prev    *DoubleNode //上节点
	next    *DoubleNode //下节点
	element Element     //存储的数据
}

//实例一个节点
func NewDoubleNode(e Element) *DoubleNode {
	return &DoubleNode{
		element: e,
	}
}

type DoubleLink struct {
	head *DoubleNode
	tail *DoubleNode
	len  int
}

func NewDoubleLink() *DoubleLink {
	return &DoubleLink{}
}

func (d *DoubleLink) Print() {
	str := ""
	p := d.head
	for p != nil {
		str += fmt.Sprint(p.element)
		p = p.next
		if p != nil {
			str += fmt.Sprint("=>")
		}
	}
	fmt.Println(str)
}

//追加节点
func (d *DoubleLink) Append(e Element) {
	node := NewDoubleNode(e)
	if d.len == 0 {
		d.head = node
		d.tail = node
	} else {
		_node := d.tail     //取出旧尾部节点
		_node.next = node   //旧的尾部点的下一个节点就是新节点
		d.tail = node       //设置新的尾部节点
		d.tail.prev = _node //设置尾部的上一个节点,即旧尾部节点
	}
	d.len++
}

//添加头部节点
func (d *DoubleLink) AddHead(e Element) {
	node := NewDoubleNode(e)
	if d.len == 0 {
		d.head = node
		d.tail = node
	} else {
		_node := d.head     //取出头部
		_node.prev = node   //将头部的下一个节点设置成新节点
		d.head = node       //将头部设置成新节点
		d.head.next = _node //让新头部的下一个节点设置成旧头部.
	}
	d.len++
}

//删除头部节点
func (d *DoubleLink) RemoveHead() Element {
	if d.len == 0 {
		return ""
	}
	node := d.head
	if node.next == nil {
		d.head, d.tail = nil, nil
	} else {
		d.head = node.next
		d.head.prev = nil
	}
	d.len--
	return node.element
}

//删除尾部节点
func (d *DoubleLink) RemoveTail() Element {
	if d.len == 0 {
		return ""
	}
	node := d.tail
	if node.prev == nil {
		d.head, d.tail = nil, nil
	} else {
		d.tail = node.prev
		d.tail.next = nil
	}
	d.len--
	return node.element
}
//func main() {
//	link := NewDoubleLink()
//	link.Append(1)
//	link.Append(2)
//	link.Append(3)
//	link.AddHead(0)
//	link.Print()
//	fmt.Println("删除头部", link.RemoveHead())
//	link.AddHead(-1)
//	link.AddHead(-2)
//	link.Print()
//	fmt.Println("删除尾部", link.RemoveTail())
//	link.Print()
//}
