package DoubleLinedList

import (
	"fmt"
	"sync"
)

//实现一个双向链表

//定义一个节点结构
//每一个节点都有一个key, value, 上节点地址, 下节点的地址
type DoubleNode struct {
	Key   int         //键
	Value interface{} //值
	Prev  *DoubleNode //上一个节点指针
	Next  *DoubleNode //下一个节点指针
	Freq  int         //频率次数.为了给LFU使用的
}

func Node(key int, value interface{}) *DoubleNode {
	return &DoubleNode{
		Key:   key,
		Value: value,
		Prev:  nil,
		Next:  nil,
	}
}

//定义一个双链表的结构
type DoubleList struct {
	lock     *sync.RWMutex //锁
	Capacity uint          //最大容量
	Size     uint          //当前容量
	Head     *DoubleNode   //头节点
	Tail     *DoubleNode   //尾部节点
}

//实现双链表

//初使双链表
func New(capacity uint) *DoubleList {
	list := new(DoubleList)
	list.Capacity = capacity
	list.lock = new(sync.RWMutex)
	list.Size = 0
	list.Head = nil
	list.Tail = nil
	return list
}

//默认向尾部添加节点
func (list *DoubleList) Append(node *DoubleNode) bool {
	return list.AddTail(node)
}

//添加头部元素
// 1. 先判断是否有容量大小
// 2. 判断头部是否为空,
// 		a. 如果为空则添加新节点
// 		b. 如果不为空则更改现有的节点,并添加上
func (list *DoubleList) AddHead(node *DoubleNode) bool {
	//判断容量是否为0
	if list.Capacity == 0 {
		return false
	}
	list.lock.Lock()
	defer list.lock.Unlock()
	//判断头部节点是否为nil
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else { //存在头部节点
		list.Head.Prev = node //将旧头部节点上一个节点指向新节点
		node.Next = list.Head //新头部节点下一个节点指向旧头部节点
		list.Head = node      //设置新的头部节点
		list.Head.Prev = nil  //将新的头部节点上一个节点设置NIL
	}
	list.Size++
	return true
}

//添加尾部元素
//追加元素
// 1. 先判断是否有容量大小
// 2. 判断尾部是否为空,
// 		a. 如果为空则添加新节点
// 		b. 如果不为空则更改现有的节点,并添加上
func (list *DoubleList) AddTail(node *DoubleNode) bool {
	//判断是否有容量,
	if list.Capacity == 0 {
		return false
	}
	list.lock.Lock()
	defer list.lock.Unlock()
	//判断尾部是否为空
	if list.Tail == nil {
		list.Tail = node
		list.Head = node
	} else {
		//旧的尾部下个节点指向新节点
		list.Tail.Next = node
		//追加新节点时,先将node的上节点指向旧的尾部节点
		node.Prev = list.Tail
		//设置新的尾部节点
		list.Tail = node
		//新的尾部下个节点设置为空
		list.Tail.Next = nil
	}
	//双链表大小+1
	list.Size++
	return true
}

//添加任意位置元素
func (list *DoubleList) Insert(index uint, node *DoubleNode) bool {
	//容量满了
	if list.Size == list.Capacity {
		return false
	}
	//如果没有节点
	if list.Size == 0 {
		return list.Append(node)
	}
	//如果插入的位置大于当前长度则尾部添加
	if index > list.Size {
		return list.AddTail(node)
	}
	//如果插入的位置等于0则,头部添加
	if index == 0 {
		return list.AddHead(node)
	}
	//取出要插入位置的节点
	nextNode := list.Get(index)
	list.lock.Lock()
	defer list.lock.Unlock()
	//中间插入则需要需要:
	//假设已有A, C节点, 现在要插入B节点
	// nextNode即是C节点,
	//A的下节点应该是B, 即C的上节点的下节点是B
	nextNode.Prev.Next = node
	//B的上节点是C的上节点
	node.Prev = nextNode.Prev
	//B的下节点是C
	node.Next = nextNode
	//C的上节点是B
	nextNode.Prev = node
	list.Size++
	return true
}

//删除任意元素
func (list *DoubleList) Remove(node *DoubleNode) *DoubleNode {
	//判断是否是头部节点
	if node == list.Head {
		return list.RemoveHead()
	}
	//判断是否是尾部节点
	if node == list.Tail {
		return list.RemoveTail()
	}
	list.lock.Lock()
	defer list.lock.Unlock()
	//节点为中间节点
	//则需要:
	//将上一个节点的下一节点指针指向下节点
	//将下一个节点的上一节点指针指向上节点
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
	return node
}

//删除头部节点
func (list *DoubleList) RemoveHead() *DoubleNode {
	//判断头部节点是否为空
	if list.Head == nil {
		return nil
	}
	list.lock.Lock()
	defer list.lock.Unlock()
	//将头部节点取出来
	node := list.Head
	//判断头部是否有下一个节点
	if node.Next != nil {
		list.Head = node.Next
		list.Head.Prev = nil
	} else { //如果没有下一个节点, 说明只有一个节点
		list.Head, list.Tail = nil, nil
	}
	list.Size--
	return node
}

//删除尾部节点
func (list *DoubleList) RemoveTail() *DoubleNode {
	//判断尾部节点是否为空
	if list.Tail == nil {
		return nil
	}
	list.lock.Lock()
	defer list.lock.Unlock()
	//取出尾部节点
	node := list.Tail
	//判断尾部节点的上一个是否存在
	if node.Prev != nil {
		list.Tail = node.Prev
		list.Tail.Next = nil
	} else {
		list.Tail, list.Head = nil, nil
	}
	list.Size--
	return node
}

//获取某一元素
func (list *DoubleList) Get(index uint) *DoubleNode {
	//如果是index = 0则返回头部
	if index == 0 {
		return list.Head
	}
	//如果超出或等于当前链大小,则返回尾部
	if index >= list.Size {
		return list.Tail
	}
	//如果中间,则需要循环index次数的链表
	var i uint
	node := list.Head
	for i = 1; i < index; i++ {
		node = node.Next
	}
	return node
}

//搜索某数据.遍历所有表
func (list *DoubleList) Search(key int) *DoubleNode {
	if list.Size == 0 {
		return nil
	}
	//从头部开始搜索
	node := list.Head
	//判断是否是头部数据.
	if node.Key == key {
		return node
	}
	//非头部数据则向下搜索
	for node != nil {
		node = node.Next
		if node.Key == key {
			return node
		}
	}
	return node
}

//获取链表大小
func (list *DoubleList) GetSize() uint {
	return list.Size
}

//打印链表
func (list *DoubleList) Print() {
	if list == nil || list.Size == 0 {
		fmt.Println("list is nil")
		return
	}
	p := list.Head
	line := ""
	for p != nil {
		line += fmt.Sprintf("key:%d, value: %v", p.Key, p.Value)
		p = p.Next
		if p != nil {
			line += fmt.Sprint(" => ")
		}
	}
	fmt.Println(line)
}

//反转链表
func (list *DoubleList) Reverse() {
	if list == nil || list.Size == 0 {
		fmt.Println("data is nil")
		return
	}
	p := list.Tail
	line := ""
	for p != nil {
		line += fmt.Sprintf("key:%d, value: %v", p.Key, p.Value)
		p = p.Prev
		if p != nil {
			line += " => "
		}
	}
	fmt.Println(line)
}
