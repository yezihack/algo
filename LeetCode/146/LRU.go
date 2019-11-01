//什么是LRU
// LRU是最近最少使用淘汰缓存策略.具体就是最近使用的放在队列的最前端,
// 包括插入,查询的,也就是只要访问过的数据都放最前端.如果容量满则淘汰最末端.

//LRU缓存机制 LeetCode 146 https://leetcode-cn.com/problems/lru-cache/solution/

//解题思路:
//首先想到的就是使用双链表实现一个队列.
//然后使用双链表的特性, 头部与尾部的特性.
//再就是借下散列函数实现O(1)取值的操作.

package main

import "fmt"

//使用双链表实现

//定义一个节点
type Node struct {
	Prev  *Node //前节点
	Next  *Node //后节点
	key   int
	value int //存储值
}

//实例一个节点
func NewNode(k, v int) *Node {
	return &Node{
		key:   k,
		value: v,
	}
}

//定义一个LRUCache结构体
type LRUCache struct {
	cache map[int]*Node //设置一个map, 用于存储值的.
	head  *Node         //头节点
	tail  *Node         //尾部节点
	len   int           //链表长度
	cap   int           //链表容量
}

func (this *LRUCache) Print() {
	str := ""
	p := this.head
	for p != nil {
		str += fmt.Sprintf("k:%d, v:%d", p.key, p.value)
		p = p.Next
		if p != nil {
			str += fmt.Sprint("=>")
		}
	}
	fmt.Println("双链表:", str)
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		cache: make(map[int]*Node), //初使化
		cap:   capacity,            //设置容量
	}
}

//添加头部节点
func (this *LRUCache) AddHead(node *Node) {
	if this.len == 0 {
		this.head = node
		this.tail = node
	} else {
		_node := this.head
		this.head = node
		_node.Prev = node
		this.head.Next = _node
	}

	this.len++
}

//需要考虑节点是否等于头部节点, 或尾部节点.都不是则是中部节点.
func (this *LRUCache) remove(n *Node) {
	if this.head == n { //判断是否是头部节点
		this.head = this.head.Next //设置新的头部节点, 即旧的头部节点的下一个节点
		this.head.Prev = nil       //设置新头部节点的上一个节点为空
	} else if this.tail == n { //判断是否是尾部节点
		this.tail = this.tail.Prev //设置尾部节点,即旧的尾部节点的上一个 .
		this.tail.Next = nil       //设置新尾部节点下一个节点为空
	} else { //此节点一定是在中间, 前有头, 后有尾
		n.Prev.Next = n.Next //将删除节点的前节点的下指针指向删除节点的下节点
		n.Next.Prev = n.Prev //将删除节点的下节点的上指针指向删除点节的前节点
	}
	delete(this.cache, n.key) //需要删除缓存里的值
	this.len--
}

func (this *LRUCache) Get(key int) int {
	//存在缓存里直接取出.并将值移动到头部.
	if node, ok := this.cache[key]; ok {
		//删除这个元素
		this.remove(node)
		//移动到头部
		this.AddHead(node)
		return node.value
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {
	//判断是否存在, 如果存在,调用get方法, 如果不存在需要判断容量是否满
	if node, ok := this.cache[key]; ok {
		this.Get(key) //调用
		//更新值
		if node.value != value {
			node.value = value
			this.cache[key] = node
		}
	} else {
		if this.len == this.cap { //如果容量满了,则淘汰最后面的数据.
			this.remove(this.tail)
		}
		node := NewNode(key, value) //实例一个节点
		this.AddHead(node)          //将数据添加到头部
		this.cache[key] = node      //将新值存入到缓存中
	}
}

func main() {
	var result int
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Print()
	result = cache.Get(1) // 返回  1
	cache.Print()
	fmt.Println(result)
	cache.Put(3, 3)
	cache.Print()
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
	cache.Print()
}
