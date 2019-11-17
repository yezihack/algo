package Interview

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"testing"
)
//查找单链表中的倒数第k个结点
func TestFindKNode(t *testing.T) {
	head := src.TestLinkedData()
	fmt.Println(PrintLinked(head))
	if p := FindKNode(1, head); p != nil {
		fmt.Println(p.Data)
	}
	if p := FindKNode(0, head); p != nil {
		fmt.Println(p.Data)
	}
	if p := FindKNode(-1, head); p != nil {
		fmt.Println(p.Data)
	}
	if p := FindKNode(10, head); p != nil {
		fmt.Println(p.Data)
	}
}
//单链表的反转
func TestReverseBySingleLinked(t *testing.T) {
	head := src.TestLinkedData()
	fmt.Println(PrintLinked(head))
	linked := ReverseBySingleLinked(head)
	fmt.Println(PrintLinked(linked))
}
//单链表的反转
func TestReverseBySingleLinked2(t *testing.T) {
	head := src.TestLinkedData()
	fmt.Println(PrintLinked(head))
	linked := ReverseBySingleLinked2(head)
	fmt.Println(PrintLinked(linked))
}
//简单反转
func TestReverse(t *testing.T) {
	linked := src.NewLinked()
	linked.Append(1)
	linked.Append(2)
	head := linked.Head
	fmt.Println(PrintLinked(head))
	next := head.Next
	head.Next = nil
	next.Next = head
	fmt.Println(PrintLinked(next))
}
//倒序打印链表, 方法一:先反转.再打印. 方法二:使用栈打印
func TestReversePrint(t *testing.T) {
	linked := src.NewLinked()
	linked.Append(1)
	linked.Append(2)
	linked.Append(3)
	fmt.Println(linked.Print())
	head := linked.Head
	ReversePrint(head)
}
//合并两个有序单链表,合并之后依然有序
func TestMergeLinked(t *testing.T) {
	l1 := src.NewLinked()
	l1.Append(1)
	l1.Append(3)
	l1.Append(5)
	h1 := l1.Head

	l2 := src.NewLinked()
	l2.Append(2)
	l2.Append(4)
	l2.Append(6)
	h2 := l2.Head

	pp := MergeLinked(h1, h2)
	fmt.Println(l1.Display(pp))
}
//合并两个有序的单链表，合并之后的链表依然有序
func TestMergeLinked2(t *testing.T) {
	l1 := src.NewLinked()
	l1.Append(1)
	l1.Append(3)
	l1.Append(5)
	h1 := l1.Head

	l2 := src.NewLinked()
	l2.Append(2)
	l2.Append(4)
	l2.Append(6)
	l2.Append(6)
	l2.Append(9)
	h2 := l2.Head

	pp := MergeLinkedByNode(h1, h2)
	fmt.Println(l1.Display(pp))
}
//链表合并,使用递归实现
func TestMergeLinkedByRecurse(t *testing.T) {
	l1 := src.NewLinked()
	l1.Append(1)
	l1.Append(3)
	l1.Append(5)
	h1 := l1.Head

	l2 := src.NewLinked()
	l2.Append(2)
	l2.Append(4)
	l2.Append(6)
	l2.Append(6)
	l2.Append(9)
	h2 := l2.Head

	pp := MergeLinkedByRecurse(h1, h2)
	fmt.Println(l1.Display(pp))
}
//判断链表是否有环 借用一个map
func TestHasCycleByMap(t *testing.T) {
	l1 := src.NewLinked()
	l1.Append(1)
	l1.Append(3)
	l1.Append(5)
	//head := l1.Head

	n1 := new(src.LinkedNode)
	n1.Data = 1

	n2 := new(src.LinkedNode)
	n2.Data = 2

	n3 := new(src.LinkedNode)
	n3.Data = 3

	n1.Next = n2
	n2.Next = n3
	n3.Next = n1
	//fmt.Println(PrintLinked(n1))

	fmt.Println(HasCycleByMap(n1))
}
//判断链表是否有环 快慢指针实现
func TestHasCycleBySlowFastPoint(t *testing.T) {
	n1 := new(src.LinkedNode)
	n1.Data = 1

	n2 := new(src.LinkedNode)
	n2.Data = 2

	n3 := new(src.LinkedNode)
	n3.Data = 3

	n1.Next = n2
	n2.Next = n3
	if HasCycleBySlowFastPoint(n1) != false {
		src.So(t, false, HasCycleBySlowFastPoint(n1))
	}
	n3.Next = n1
	if HasCycleBySlowFastPoint(n1) != true {
		src.So(t, true, HasCycleBySlowFastPoint(n1))
	}
}
func TestSwapPairs(t *testing.T) {
	l1 := src.NewLinked()
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	l1.Append(4)
	l1.Append(5)
	head := l1.Head
	fmt.Println(PrintLinked(SwapPairs(head)))
}

func TestSwapPairsByReserve(t *testing.T) {
	l1 := src.NewLinked()
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	l1.Append(4)
	l1.Append(5)
	head := l1.Head
	fmt.Println(PrintLinked(SwapPairsByReserve(head)))
}
func TestRLinked(t *testing.T) {
	l1 := src.NewLinked()
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	l1.Append(4)
	l1.Append(5)
	head := l1.Head
	fmt.Println(PrintLinked(RLinked(head)))
}