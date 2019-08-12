# 双链表的实现

## 基本概念
> 每一个节点都存储上一个和下一个节点的指针

## 实现思路

>  创建一个节点结构体
1. 每个节点都有上节点指针与下节点指针
1. 每个节点都有一个key => value

> 创建一个链表结构体
1. 链表容量大小属性
1. 链表大小属性
1. 链表锁, 实现并发安全
1. 链表头节点
1. 链表尾节点

> 实现链表操作方法
1. 添加头部节点操作AppendHead
1. 添加尾部节点操作AppendTail
1. 追加尾部节点操作Append
1. 插入任意节点操作Insert
1. 删除任意节点操作Remove
1. 删除头部节点操作RemoveHead
1. 删除尾部节点操作RemoveTail
1. 获取指定位置的节点Get
1. 搜索任意节点Search
1. 获取链表大小GetSize
1. 打印所有节点操作Print
1. 反转所有节点操作Reverse

## 代码解析

### 定义节点的结构体
```
type DoubleNode struct {
	Key   int         //键
	Value interface{} //值
	Prev  *DoubleNode //上一个节点指针
	Next  *DoubleNode //下一个节点指针
	Freq  int         //频率次数.为了给LFU使用的
}
```


### 定义一个双链表的结构体
```
//定义一个双链表的结构
type DoubleList struct {
	lock     *sync.RWMutex //锁
	Capacity uint          //最大容量
	Size     uint          //当前容量
	Head     *DoubleNode   //头节点
	Tail     *DoubleNode   //尾部节点
}
```


### 初使双链表
```
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
```

### 添加头部节点
> 实现思路
1. 先判断容量大小
1. 判断头部是否为空,
    1. 如果为空则添加新节点
    1. 如果不为空则更改现有的节点,并添加上

```
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
```

### 添加尾部元素
> 实现思路
1. 先判断容量大小
1. 判断尾部是否为空,
    1. 如果为空则添加新节点
    1. 如果不为空则更改现有的节点,并添加上

```

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
```

### 添加任意位置元素
> 实现思路
1. 判断容量大小
1. 判断链表大小
1. 第一种: 如果插入的位置大于当前长度则尾部添加
1. 第二种: 如果插入的位置等于0则,头部添加
1. 第三种: 中间插入节点
    1. 先取出要插入位置的节点,假调为C节点
    1. 介于A, C之间, 插入一个B节点
    1. A的下节点应该是B, 即C的上节点的下节点是B
    1. B的上节点是C的上节点
    1. B的下节点是C

```
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
	//中间插入:
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
```

### 删除头部节点
> 实现思路
1. 判断头部是否为空
1. 将头部节点取出来
1. 判断头部是否有下一个节点
    1. 没有下一个节点,则说明只有一个节点, 删除本身,无须移动指针位置
    1. 如果有下一个节点,则头部下一个节点即是头部节点.
```
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
```

### 删除尾部节点
> 实现思路
1. 判断尾部节点是否为空
1. 取出尾部节点
1. 判断尾部节点的上一个节点是否存在
    1. 不存在:则说明只有一个节点, 删除本身,无须移动指针位置
    1. 如果存在上一个节点,则尾部的上一个节点即是尾部节点.
```
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
	}
	list.Size--
	return node
}
```

### 删除任意元素
> 实现思路
1. 判断是否是头部节点
1. 判断是否是尾部节点
1. 否则为中间节点,需要移动上下节点的指针位置.并实现元素删除
    1. 将上一个节点的下一节点指针指向下节点
    1. 将下一个节点的上一节点指针指向上节点
```
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
```

## [查看完整源码](https://github.com/yezihack/math/blob/master/DoubleLinedList/double_lined_list.go)




