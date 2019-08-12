package Cache

import (
	"github.com/yezihack/math/DoubleLinedList"
)

/* FIFO(first in first out) 先进先出淘汰算法*/

/*
解题思路:
	1. 借用双链表实现FIFO
	2. 借用map实现缓存, 避免对链表的搜索
*/

//定义FIFO的结构体
type FIFOCache struct {
	double *DoubleLinedList.DoubleList
	cache  map[int]*DoubleLinedList.DoubleNode
}

//实例对象
func NewFIFO(capacity uint) *FIFOCache {
	fifo := FIFOCache{}
	fifo.double = DoubleLinedList.New(capacity)
	fifo.cache = make(map[int]*DoubleLinedList.DoubleNode)
	return &fifo
}

//获取值
func (f *FIFOCache) Get(key int) *DoubleLinedList.DoubleNode {
	//判断是否在缓存中
	if node, ok := f.cache[key]; ok { //存在的话
		return node
	} else { //不存在的话
		return nil
	}
}

//添加元素
func (f *FIFOCache) Put(key int, data interface{}) bool {
	//判断是否存在缓存中
	if node, ok := f.cache[key]; ok { //存在的话
		//更新数据,删除链表, 再添加到链表中
		node.Value = data
		f.double.Remove(node)
		f.double.Append(node)
	} else { //不存在的话
		////判断容量是否满了
		//if f.double.Size >= f.double.Capacity { //满了.
		//	//删除头部数据,因为我们是追加, 头部数据是最早添加的
		//	f.double.RemoveHead()
		//	node = DoubleLinedList.Node(key, data)
		//	f.double.Append(node)
		//	f.cache[key] = node
		//} else { //没满
		//	node = DoubleLinedList.Node(key, data)
		//	f.double.Append(node)
		//	f.cache[key] = node
		//}
		/**************将以上逻辑简化*****************/
		if f.double.Size >= f.double.Capacity { //满了.
			//删除头部数据,因为我们是追加, 头部数据是最早添加的
			f.double.RemoveHead()
		}
		node = DoubleLinedList.Node(key, data)
		f.double.Append(node)
		f.cache[key] = node
	}
	return true
}

//获取链表大小
func (f *FIFOCache) GetSize() uint {
	return f.double.Size
}

//打印
func (f *FIFOCache) Print() {
	f.double.Print()
}
