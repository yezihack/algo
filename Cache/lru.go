package Cache

import (
	"github.com/yezihack/math/DoubleLinedList"
)

/* LRU(least recently used) 最近最少使用淘汰算法 */
/*
	实现思路:
	1.使用到的数据, 放头部
	2.如果超出容量则将最尾部淘汰掉
	3.设置一个map用于避免使用搜索链表, 因为遍历链表速度很慢
*/

//定义最近最小使用结构体
type LRUCache struct {
	double *DoubleLinedList.DoubleList         //使用双向链表
	cache  map[int]*DoubleLinedList.DoubleNode //用于存储使用过的值
}

//实例LRU对象
func NewLRUCache(capacity uint) *LRUCache {
	l := &LRUCache{}
	l.double = DoubleLinedList.New(capacity)
	l.cache = make(map[int]*DoubleLinedList.DoubleNode)
	return l
}

//使用数据.
// 1. 将使用过的数据,添加到头部.
// 2. 当容量满时,则删除尾部数据.
func (l *LRUCache) Get(key int) *DoubleLinedList.DoubleNode {
	//判断缓存里是否存在
	if node, ok := l.cache[key]; ok {
		//1.删除节点, 2.添加到头部
		l.double.Remove(node)
		l.double.AddHead(node)
		return node
	} else {
		return nil
	}
}

//添加数据
func (l *LRUCache) Put(key int, data interface{}) bool {
	if node, ok := l.cache[key]; ok { //key存在缓存中
		//1. 更新数据.
		node.Value = data
		//2. 移除之前的链表数据.
		l.double.Remove(node)
		//3. 添加到头部链表数据.
		l.double.AddHead(node)
		l.cache[key] = node
	} else { //key不存在缓中
		//判断容量是否满了
		// 1.容量满了.
		if l.double.Size >= l.double.Capacity {
			//a. 淘汰尾部数据,b.添加头部数据, c.添加缓存
			l.double.RemoveTail()
		}
		//代码合并
		node = DoubleLinedList.Node(key, data)
		l.double.AddHead(node)
		l.cache[key] = node
	}
	return true
}

func (l *LRUCache) GetHead() *DoubleLinedList.DoubleNode {
	return l.double.Get(0)
}

//获取当前大小
func (l *LRUCache) GetSize() uint {
	return l.double.Size
}

//打印数据
func (l *LRUCache) Print() {
	l.double.Print()
}
