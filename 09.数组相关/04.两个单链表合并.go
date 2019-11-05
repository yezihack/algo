package main

import (
	"fmt"
	. "github.com/yezihack/algo/09.数组相关/src"
)

func main() {
	s1 := NewSingleLink()
	s1.Append(NewSingleNode(1))
	s1.Append(NewSingleNode(7))
	s1.Append(NewSingleNode(8))
	s1.Print()
	s2 := NewSingleLink()
	s2.Append(NewSingleNode(2))
	s2.Append(NewSingleNode(4))
	s2.Append(NewSingleNode(6))
	s2.Append(NewSingleNode(8))
	s2.Append(NewSingleNode(10))
	s2.Print()

	node := MergeLinkedByFor(s1.Head, s2.Head)
	Print("result", node)

}
func MergeLinkedByFor(n1, n2 *SingleNode) *SingleNode {
	p1, p2 := n1, n2              //创建二个指针
	var pp, pp2, next *SingleNode //创建一个起启指针PP, 再创建一个移动的指针 PP2, 再创建一个NEXT指针指向最小值的链上
	for p1 != nil && p2 != nil {
		if p1.Data <= p2.Data {
			next = p1.Next
			if pp == nil {
				pp = p1
			} else {
				pp2.Next = p1
			}
			p1.Next = nil
			pp2 = p1
			p1 = next
		} else {
			next = p2.Next
			if pp == nil {
				pp = p2
			} else {
				pp2.Next = p2
			}
			p2.Next = nil
			pp2 = p2
			p2 = next
		}
	}
	return pp
}

//for实现链表合并
//参考: https://blog.csdn.net/meng_lemon/article/details/82080454
func MergeLinked(n1, n2 *SingleNode) *SingleNode {
	p1, p2 := n1, n2
	var pp, next, tail *SingleNode
	for p1 != nil && p2 != nil {
		if p1.Data < p2.Data {
			next = p1.Next
			if pp == nil {
				pp = p1
			} else {
				tail.Next = p1
			}
			p1.Next = nil
			tail = p1
			p1 = next
		} else {
			next = p2.Next
			if pp == nil {
				pp = p2
			} else {
				tail.Next = p2
			}
			p2.Next = nil
			tail = p2
			p2 = next
		}
	}
	Print("p1", p1)
	Print("p1", p2)
	if p1 != nil {
		tail.Next = p1
	}
	if p2 != nil {
		tail.Next = p2
	}
	return pp
}

//递归实现链表合并
func MergeLinked2(n1, n2 *SingleNode) *SingleNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Data <= n2.Data {
		n1.Next = MergeLinked2(n1.Next, n2)
		return n1
	} else {
		n2.Next = MergeLinked2(n1, n2.Next)
		return n2
	}
}

//递归实现链表合并
func LinkMerge(n1, n2 *SingleNode) *SingleNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	var n3 *SingleNode
	//如果n1值小于n2的值, 则将n1指针指向n3,然后继续下一个n1的结点比较
	if n1.Data <= n2.Data {
		n3 = n1
		n3.Next = LinkMerge(n1.Next, n2)
	} else {
		n3 = n2
		n3.Next = LinkMerge(n1, n2.Next)
	}
	return n3
}

func Print(flag string, node *SingleNode) {
	str := flag + ":\n"
	for node != nil {
		str += fmt.Sprint(node.Data)
		node = node.Next
		if node != nil {
			str += "=>"
		}
	}
	fmt.Println(str)
}
