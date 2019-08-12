package DoubleLinedList

import (
	"testing"
)

func TestNew(t *testing.T) {
	list := New(10)
	nodes := make([]*DoubleNode, 0)
	for i := 0; i < 10; i++ {
		node := Node(i, i)
		nodes = append(nodes, node)
	}
	list.Append(nodes[0])
	list.Print()
	list.Append(nodes[1])
	list.Print()
	list.Append(nodes[2])
	list.Print()
	list.RemoveTail()
	list.Print()
	list.RemoveHead()
	list.Print()
	list.Append(nodes[3])
	list.Print()
	list.Insert(2, nodes[4])
	list.Print()
	list.Append(nodes[5])
	list.Print()
	list.Append(nodes[6])
	list.Print()
	list.Reverse()
}
func TestDoubleList_Print(t *testing.T) {
	l := New(5)
	for i := 0; i < 5; i++ {
		node := Node(i, i)
		l.Append(node)
	}
	if l.GetSize() != 5 {
		t.Errorf("want: %d, errValue: %d", 5, l.GetSize())
	}
	l.Print()
}
