package queue_10

type LinkedNode struct {
	front *LinkedNode //前趋指针
	rear  *LinkedNode //后续指针
	data  interface{} //存储的值
}

type LinkedQueue struct {
	Head *LinkedNode //头指针
	Tail *LinkedNode //尾指针
}
