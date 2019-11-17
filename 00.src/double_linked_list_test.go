package src

import (
	"fmt"
	"testing"
)
func TestNewDLinkedList(t *testing.T) {
	l := NewDLinkedList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	fmt.Println(l.Print())
	var expect interface{}
	expect = "1=>2=>3=>4"
	if l.Print() != expect {
		t.Errorf("expect:%s, actual:%s\n", expect, l.Print())
	}
}
func TestDLinkedList_Reverse(t *testing.T) {
	l := NewDLinkedList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	fmt.Println(l.Print())
	head := l.Reverse()
	fmt.Println(l.Display(head))
	var expect interface{}
	expect = 4
	if head.Data != expect {
		t.Errorf("expect:%d, actual:%d\n", expect, head.Data)
	}
	expect = 3
	if head.Next.Data != expect {
		t.Errorf("expect:%d, actual:%d\n", expect, head.Next.Data)
	}
}
func TestDLinkedList_RemoveHead(t *testing.T) {
	l := NewDLinkedList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	fmt.Println(l.Print())
	So(t, 1, l.RemoveHead())
	fmt.Println(l.Print())
	So(t, 2, l.RemoveHead())
	fmt.Println(l.Print())
	So(t, 3, l.RemoveHead())
	fmt.Println(l.Print())
	So(t, 4, l.RemoveHead())
	fmt.Println(l.Print())
	So(t, nil, l.RemoveHead())
	fmt.Println(l.Print())
}
func TestDLinkedList_RemoveTail(t *testing.T) {
	l := NewDLinkedList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	fmt.Println(l.Print())
	So(t, 4, l.RemoveTail())
	fmt.Println(l.Print())
	So(t, 3, l.RemoveTail())
	fmt.Println(l.Print())
	So(t, 2, l.RemoveTail())
	fmt.Println(l.Print())
	So(t, 1, l.RemoveTail())
	fmt.Println(l.Print())
}
func TestDLinkedList_Remove(t *testing.T) {
	l := NewDLinkedList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	fmt.Println(l.Print())
	So(t, 4, l.Remove(4))
	fmt.Println(l.Print())
	So(t, 1, l.Remove(1))
	fmt.Println(l.Print())
	So(t, 2, l.Remove(1))
	fmt.Println(l.Print())
	So(t, 3, l.Remove(1))
	fmt.Println(l.Print())
	fmt.Println(l.Print())
}