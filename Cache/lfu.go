package Cache

import (
	"fmt"
	"github.com/yezihack/math/DoubleLinedList"
	"math"
)

/* LFU(least frequently used) 最不常用使用淘汰算法 */
/*
解题思路:
	1. 需要记录使用次数,相同的使用次数记录在一起
	2. 添加, 获取都记录为一次.
	3. 当容量满时,淘汰次数最少的链表
*/

//定义一个切片链表
type FreqList struct {
	list []*DoubleLinedList.DoubleNode
}

//定义一个最不常用的结构体
type LFUCache struct {
	cache     map[int]*DoubleLinedList.DoubleNode //缓存节点
	cacheFreq map[int]*DoubleLinedList.DoubleList //key:为频率, value:为频率链表
	capacity  int                                 //容量大小
	size      int                                 //当前大小
}

func NewLFUCache(capacity int) *LFUCache {
	f := LFUCache{}
	f.capacity = capacity
	f.cache = make(map[int]*DoubleLinedList.DoubleNode)
	f.cacheFreq = make(map[int]*DoubleLinedList.DoubleList, 0)
	return &f
}

//更新频率
func (f *LFUCache) updateFreq(node *DoubleLinedList.DoubleNode) {
	// 删除频率MAP里的节点,再加+1, 再存入
	freq := node.Freq
	//删除频率里的MAP节点
	node = f.cacheFreq[freq].Remove(node)
	if f.cacheFreq[freq].Size == 0 { //如果这个频率里的链表大小为0则删除掉
		delete(f.cacheFreq, freq)
	}
	//更新, .
	freq++
	node.Freq = freq
	//新的频率了, 需要判断是否在缓存中
	if _, ok := f.cacheFreq[freq]; !ok { //不存在缓存中, 则创建新的.
		f.cacheFreq[freq] = DoubleLinedList.New(math.MaxInt32)
	}
	f.cacheFreq[freq].Append(node)
}

//获取
func (f *LFUCache) Get(key int) *DoubleLinedList.DoubleNode {
	if node, ok := f.cache[key]; ok {
		f.updateFreq(node)
		return node
	}
	return nil
}

//使用
func (f *LFUCache) Put(key int, data interface{}) bool {
	if f.capacity == 0 {
		return false
	}
	//判断是否在缓存中
	if node, ok := f.cache[key]; ok { //缓存里有
		node.Value = data
		f.updateFreq(node)
		return true
	} else { //缓存里无
		if f.size == f.capacity { //满了
			//淘汰掉频率最低的数据.
			freqKey := f.MinFreq()
			node = f.cacheFreq[freqKey].Head
			delete(f.cache, node.Key)
			delete(f.cacheFreq, freqKey)
			f.size--
		}
		//添加数据.
		node = DoubleLinedList.Node(key, data)
		node.Freq = 1
		if _, ok := f.cacheFreq[node.Freq]; !ok {
			f.cacheFreq[node.Freq] = DoubleLinedList.New(math.MaxUint)
		}
		f.cacheFreq[node.Freq].Append(node)
		f.cache[key] = node
		f.size++
	}
	return true
}

//求出MAP里最小值的下标
func (f *LFUCache) MinFreq() int {
	min := math.MaxInt
	for freqKey := range f.cacheFreq {
		if freqKey < min {
			min = freqKey
		}
	}
	return min
}

//获取最大频率下标
func (f *LFUCache) MaxFreq() int {
	max := 0
	for freqKey := range f.cacheFreq {
		if freqKey > max {
			max = freqKey
		}
	}
	return max
}

//打印
func (f *LFUCache) Print() {
	fmt.Println("---------------------------")
	for k := range f.cacheFreq {
		fmt.Printf("freq:%d \n", k)
		f.cacheFreq[k].Print()
	}
	fmt.Println("---------------------------")
}
