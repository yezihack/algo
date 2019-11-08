package main

//定义一个循环队列
type QueueRaverse struct {
	data  []*BiNode
	cap   int //队列最大存储数量
	len   int //当前队列数量
	front int //前指针,前指针负责添加数据移动.
	rear  int //尾指针,后指针负责弹出数据移动.
}

//init queue
func NewQueue(cap int) *QueueRaverse {
	return &QueueRaverse{
		data: make([]*BiNode, cap),
		cap:  cap,
	}
}
func (q *QueueRaverse) Push(data *BiNode) {
	if q.QueueFull() {
		q.len--
	}

}

func (q *QueueRaverse) QueueEmpty() bool {
	if q.len == 0 {
		return true
	}
	return false
}
func (q *QueueRaverse) QueueFull() bool {
	if q.len == q.cap {
		return true
	}
	return false
}
