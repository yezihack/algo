package linked

import (
	"fmt"
	"strconv"
	"testing"
)

type link struct {
	val int
	next *link
}
//反转链表思路:
//头插法.
func ReversalLink(head *link) *link {
	var prev *link //创建一个空链表.
	for head != nil {
		curr := head.next //获取链表下一个节点
		head.next = prev //将链表断开节点.
		prev = head //链表断开的前部分给新的链表
		head = curr //断开的后部分给之前的链表表.
	}
	return prev
}

func TestReversalLink(t *testing.T) {
	l0 := &link{val:1}
	l1 := &link{val:2}
	l2 := &link{val:3}
	l0.next = l1
	l1.next = l2
	head := l0
	head = ReversalLink(head)
	str := ""
	for head != nil {
		str += strconv.Itoa(head.val)
		head = head.next
		if head != nil {
			str += "=>"
		}
	}
	fmt.Println(str)
}
