package main

import "fmt"

//循环队列使用线性表表示与实现
//如何解决队空还是队满有三种办法
//1.使用一个变量记录入栈时就加1, 出栈时就减1
//2.另设一个标记区别队空还是队满
//3.少用一个元素. <本代码就是使用此方法,见IsEmpty,IsFull>
type CycleQueue struct {
	data  []int //存储空间
	front int   //前指针,前指针负责添加数据移动.
	rear  int   //尾指针,后指针负责弹出数据移动.
	cap   int   //设置切片最大容量
	len   int   //记录元素的个数
}

//init Queue
func NewCycleQueue(cap int) *CycleQueue {
	return &CycleQueue{
		data: make([]int, cap),
		cap:  cap,
	}
}

//因为是循环队列, 后指针减去前指针 加上最大值, 然后与最大值 取余
func (q *CycleQueue) QueueLength() int {
	fmt.Printf("front:%d, rear:%d, cap:%d\n", q.front, q.rear, q.cap)
	return (q.front - q.rear + q.cap) % q.cap
}

//记录真实长度
func (q *CycleQueue) Len() int {
	fmt.Printf("front:%d, rear:%d, cap:%d\n", q.front, q.rear, q.cap)
	return q.len
}

//入队操作
//需要操作就是队满
func (q *CycleQueue) Push(n int) {
	if q.IsFull() {
		q.len--
	}
	q.data[q.rear] = n
	q.rear = (q.rear + 1) % q.cap
	q.len++
	q.IsFull()
}

//出队操作
//需要考虑: 队空没有数据返回了
func (q *CycleQueue) Pop() int {
	if q.IsEmpty() {
		return -1
	}
	t := q.data[q.front]
	q.data[q.front] = 0
	q.front = (q.front + 1) % q.cap //因为是循环队列, 指针指到数组下标最大值+1状态下
	//相当于数组越界了, 所有我们需要让指针回到起始位置,也就是除以最大值取余即可.
	q.len--
	return t
}
func (q *CycleQueue) Remote() {

}

//判断是否为空, 也就是二个指针相遇了,就表示没有数组啦.
func (q *CycleQueue) IsEmpty() bool {
	if q.front == q.rear && q.len == 0 {
		return true
	}
	return false
}

//判断队列是否满了
func (q *CycleQueue) IsFull() bool {
	if q.len == q.cap {
		return true
	}
	return false
}

//查看所有元素
func (q *CycleQueue) Display() string {
	return fmt.Sprintf("%v", q.data)
}
func main() {
	c := NewCycleQueue(3)
	c.Push(1)
	c.Push(2)
	//c.Push(3)
	fmt.Println("realLength", c.Len(), "queue", c.Display())
	fmt.Println(c.Pop())
	fmt.Println("realLength", c.Len(), "queue", c.Display())
	c.Push(9)
	fmt.Println("realLength", c.Len(), "queue", c.Display())
	c.Push(10)
	fmt.Println("realLength", c.Len(), "queue", c.Display())
	c.Push(11)
	fmt.Println("realLength", c.Len(), "queue", c.Display())

}
