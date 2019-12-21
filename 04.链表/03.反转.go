package linked

import (
	"fmt"
	. "github.com/yezihack/algo/00.src"
)

//反转打印链表, 使用头插法
func ReverseLinked(head *LinkedNode) *LinkedNode {
	var curr *LinkedNode //定义一个当前节点
	for head != nil {
		next := head.Next //1.下一个节点
		head.Next = curr//3.头的下一个节点设置为当前节点
		curr = head //4.当前节点设置为头节点
		head = next //2.将头节点移动下一个节点
	}
	return curr
}
//递归反转链表
func ReserveLinkedRecursion(head *SingleNode) *SingleNode {
	if head == nil || head.next == nil {
		return head
	}
	newHead := ReserveLinkedRecursion(head.next)
	head.next.next = head
	head.next = nil
	return newHead

}

//反转链表,非递归方法
//时间复杂度：O(n)O(n)，假设 nn 是列表的长度，时间复杂度是 O(n)O(n)。
//空间复杂度：O(1)O(1)。
func Reverse(head *SingleNode) *SingleNode {
	if head == nil || head.next == nil{
		return nil
	}
	//申请一个前驱指针
	var pre *SingleNode
	//设置一个当前指针
	cur := head.next
	for cur != nil { //
		tmp := cur.next //1 定义一个临时指针,指向下一个指针
		cur.next = pre //3 当前的下一个指向前驱指针
		pre = cur //4 前驱指针指向当前
		cur = tmp //2 当前指向下一个指针
	}
	head.next = pre //前驱指向头的下一个结点
	return head
}

//反转链表
func Inversion(head *LinkedNode) *LinkedNode {
	var prev *LinkedNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
//头插法
func InversionByHead(head *LinkedNode) *LinkedNode {
	var newHead *LinkedNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}

//打印链表
func PrintLinked(head *SingleNode) {
	p := head.next
	format := ""
	for p != nil {
		format += fmt.Sprint(p.GetData())
		p = p.next
		if p != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
