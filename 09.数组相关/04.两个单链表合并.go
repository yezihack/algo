package main

import (
	"fmt"
	. "github.com/yezihack/math/09.数组相关/src"
	"runtime/debug"
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

	node := MergeLinked2(s1.Head, s2.Head)
	Print("result", node)

}
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

func mergeTwoLists(l1 *SingleNode, l2 *SingleNode) *SingleNode {
	fmt.Printf("%s", debug.Stack())
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Data <= l2.Data {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

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
