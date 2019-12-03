package queue_10

import "fmt"

//https://www.bilibili.com/video/av33876282
//循环队列使用线性表表示与实现
//如何解决队空还是队满有三种办法
//1.使用一个变量记录入栈时就加1, 出栈时就减1
//2.另设一个标记区别队空还是队满
//3.少用一个元素. <本代码就是使用此方法,见IsEmpty,IsFull>
//4.队列满了则不允许添加啦
//5.少用一个元素,当rear+1就表示满了.
type CycleQueue struct {
	data  []interface{} //存储空间
	front int           //前指针,前指针负责弹出数据移动
	rear  int           //尾指针,后指针负责添加数据移动
	cap   int           //设置切片最大容量
}

//init Queue
func NewCycleQueue(cap int) *CycleQueue {
	return &CycleQueue{
		data:  make([]interface{}, cap),
		cap:   cap,
		front: 0,
		rear:  0,
	}
}

//因为是循环队列, 后指针减去前指针 加上最大值, 然后与最大值 取余
func (q *CycleQueue) QueueLength() int {
	return (q.rear - q.front + q.cap) % q.cap
}

//入队操作
//判断队列是否队满,队满则不允许添加数据
func (q *CycleQueue) Push(data interface{}) bool {
	//check queue is full
	if (q.rear+1)%q.cap == q.front { //队列已满时，不执行入队操作
		return false
	}
	q.data[q.rear] = data         //将元素放入队列尾部
	q.rear = (q.rear + 1) % q.cap //尾部元素指向下一个空间位置,取模运算保证了索引不越界（余数一定小于除数）
	return true
}

//出队操作
//需要考虑: 队空没有数据返回了
func (q *CycleQueue) Pop() interface{} {
	if q.rear == q.front {
		return nil
	}
	data := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % q.cap
	return data
}

//判断队列是否满了
func (q *CycleQueue) QueueFull() bool {
	if (q.rear+1)%q.cap == q.front {
		return true
	}
	return false
}
func (q *CycleQueue) QueueEmpty() bool {
	if q.front == q.rear {
		return true
	}
	return false
}

//查看所有元素
func (q *CycleQueue) Display() string {
	return fmt.Sprintf("%v", q.data)
}
//func main() {
//	c := NewCycleQueue(3)
//	fmt.Println(c.Push(1))
//	fmt.Println(c.Push(2))
//	fmt.Println(c.Push(3))
//	fmt.Println(c.Push(4))
//	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
//	fmt.Println(c.Pop())
//	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
//	c.Push(9)
//	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
//	fmt.Println(c.Pop())
//	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
//	c.Push(11)
//	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
//	fmt.Println(c.Pop())
//	fmt.Println(c.Pop())
//	fmt.Println("realLength", c.QueueLength(), "queue", c.Display())
//	fmt.Println(c.QueueFull(), c.QueueEmpty())
//}
