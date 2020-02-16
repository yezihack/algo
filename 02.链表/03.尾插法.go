package linked


//defined linked struct
type linkedV2 struct {
	data int //value store
	next *linkedV2 // next linked node
}
//尾插法.
func TailInsertLinked(arr []int) *linkedV2 {
	var head = new(linkedV2)		//创建一个头节点.
	tail := head					//赋值给一个尾节点变量.
	for i := 0;i < len(arr);i ++ {
		node := &linkedV2{data:arr[i]}//创建一个节点
		tail.next = node			  //接在尾节点的下一个节点上.
		tail = node					  //设置新的尾节点.
	}
	return head.next				  //因为之前创建了一个头节点,并没有存储数据.只是为了方便操作, 所以返回头节点下一个节点.
}
