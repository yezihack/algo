package Interview

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
)

//链表相关

//打印链表.方便调试
func PrintLinked(head *src.LinkedNode) string {
	str := ""
	p := head
	for p != nil {
		str += fmt.Sprint(p.Data)
		if p.Next != nil {
			str += "=>"
		}
		p = p.Next
	}
	return str
}

//查找单链表中的倒数第k个结点 【新浪面试题】
func FindKNode(k int, head *src.LinkedNode) *src.LinkedNode {
	size := 0
	//先获取链表总长度
	p := head
	for p != nil {
		size ++
		p = p.Next
	}
	//判断是否越界
	if k <= 0 || k > size {
		return nil
	}
	//计算正序的位置,总长度减去倒数第k的位置
	pp := head
	for i := 0; i < (size - k); i++ {
		pp = pp.Next
	}
	return pp
}
//单链表的反转【腾讯面试题，有点难度】
func ReverseBySingleLinked(head *src.LinkedNode) *src.LinkedNode {
	//如果结点为空或只有一个结点时,无需反转
	if head == nil && head.Next == nil {
		return nil
	}
	cur, prev := head, new(src.LinkedNode)
	for cur != nil {
		cur.Next, cur, prev = prev, cur.Next, cur
	}
	return prev
}
//单链表的反转2
func ReverseBySingleLinked2(head *src.LinkedNode) *src.LinkedNode {
	//如果结点为空或只有一个结点时,无需反转
	if head == nil && head.Next == nil {
		return nil
	}
	//声明两个空指针结点,pre为前趋结点, next为后趋结点
	var pre, next *src.LinkedNode
	for head != nil {
		//1.将head.next赋值给next变量，也就是说next指向了节点2，先将节点2保存起来。
		next = head.Next
		//2.将pre变量赋值给了head.next，即节点1指向了null。
		head.Next = pre
		//3.将head赋值给了pre，即pre指向节点1，将节点1设为“上一个节点”。
		pre = head
		//4.将next赋值给head，即head指向了节点2。将节点2设置为“头结点”。
		head = next
	}
	return pre
}
//从尾到头打印单链表 【百度，要求方式1：反向遍历 。 方式2：Stack栈】
func ReversePrint(head *src.LinkedNode) {
	s := src.NewStack()
	for head != nil {
		s.Push(head.Data)
		head = head.Next
	}
	size := s.Length()
	for i := 0; i < size; i ++ {
		fmt.Print(s.Pop(), "=>")
	}
}
//合并两个有序的单链表，合并之后的链表依然有序
func MergeLinked(head1, head2 *src.LinkedNode) *src.LinkedNode {
	l := src.NewLinked()
	for head1 != nil && head2 != nil {
		if head1.Data.(int) < head2.Data.(int) {
			l.Append(head1.Data)
			head1 = head1.Next
		} else {
			l.Append(head2.Data)
			head2 = head2.Next
		}
	}
	if head1 != nil {
		l.Append(head1.Data)
		head1 = head1.Next
	}
	for head2 != nil {
		l.Append(head2.Data)
		head2 = head2.Next
	}
	return l.Head
}
//合并两个有序的单链表，合并之后的链表依然有序
func MergeLinkedByNode(h1, h2 *src.LinkedNode) *src.LinkedNode {
	head := new(src.LinkedNode)
	node := head
	for h1 != nil && h2 != nil {
		if h1.Data.(int) > h2.Data.(int) {
			head.Next = h2
			//head = h2
			h2 = h2.Next
		} else {
			head.Next = h1
			//head = h1
			h1 = h1.Next
		}
		head = head.Next
	}
	if h1 != nil {
		head.Next = h1
	}
	if h2 != nil {
		head.Next = h2
	}
	return node.Next
}
//链表合并,使用递归实现
//递归的核心方法是将问题规模不断缩小化
// 合并两个长度为n和m的链表，在value(n) < value(m)可以将规模缩减为合并长度为(n-1)和m的链表
func MergeLinkedByRecurse(h1, h2 *src.LinkedNode) *src.LinkedNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if h1.Data.(int) < h2.Data.(int) {
		h1.Next = MergeLinkedByRecurse(h1.Next, h2)
		return h1
	} else {
		h2.Next = MergeLinkedByRecurse(h1, h2.Next)
		return h2
	}
}


//判断链表是否有环 借用一个map
func HasCycleByMap(head *src.LinkedNode) bool {
	set := make(map[interface{}]struct{})
	for head != nil {
		if _, ok := set[head.Data]; ok {
			return true
		}
		set[head.Data] = struct{}{}
		head = head.Next
	}
	return false
}
//判断链表是否有环 快慢指针实现
func HasCycleBySlowFastPoint(head *src.LinkedNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
//链表交换相邻元素
//https://www.jianshu.com/p/68e215a129bd
func SwapPairs(head *src.LinkedNode) *src.LinkedNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	prev := new(src.LinkedNode)
	prev.Next = head
	p := prev
	for p.Next != nil && p.Next.Next != nil {
		start := p.Next
		end := p.Next.Next
		p.Next = end
		start.Next = end.Next
		end.Next = start
		p = start
	}
	return prev.Next
}
//递归实现,相邻元素交换
func SwapPairsByReserve(head *src.LinkedNode) *src.LinkedNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := head.Next
	head.Next = SwapPairsByReserve(head.Next.Next)
	n.Next = head
	return n
}
func SwapLinked(head *src.LinkedNode) *src.LinkedNode {
	if head == nil || head.Next == nil {
		return nil
	}
	n := head.Next
	head.Next = SwapLinked(head.Next.Next)
	n.Next = head
	return n
}

func RLinked(head *src.LinkedNode) *src.LinkedNode {
	var prev, next *src.LinkedNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev.Next
}




















