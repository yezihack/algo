package linked

//判断链表是否有环,使用快慢指针
func HasCycle(head *SingleNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for slow != nil && fast != nil {
		slow = slow.next//慢指针走一步
		fast = fast.next.next//快指针走二步
		if slow == fast {
			return true
		}
	}
	return false
}
