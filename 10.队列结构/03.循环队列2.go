package queue_10

import "fmt"

//循环队列,使用二个指针.入队列移动tail, 出队列移动head
//重点:判断队空与队满.利用空余空间
type CycleQ struct {
	items []interface{}
	head int
	tail int
	cap int
}
func NewCycleQ(capacity int) *CycleQ {
	return &CycleQ{
		items:make([]interface{}, capacity),
		cap: capacity,
	}
}
//入队操作
func (q *CycleQ) Push(v interface{}) bool {
	//判断队满否
	if (q.tail + 1) % q.cap == q.head {
		return false
	}
	q.items[q.tail] = v
	q.tail = (q.tail + 1) % q.cap
	return true
}
//出队操作
func (q *CycleQ) Pop() interface{} {
	//判断是否队空
	if q.tail == q.head {
		return nil
	}
	v := q.items[q.head]
	q.head = (q.head + 1) % q.cap
	return v
}
//队列元素
func (q *CycleQ) Length() int {
	fmt.Printf("tail:%d, head:%d\n", q.tail, q.head)
	return (q.tail - q.head + q.cap) % q.cap
}
