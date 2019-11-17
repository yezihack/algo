package src

import (
	"fmt"
	"testing"
)

func TestNewLinked(t *testing.T) {
	l := NewLinked()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	fmt.Println(l.Print())
}
func TestLinked_Reserve(t *testing.T) {
	var expect interface{}
	l := NewLinked()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	fmt.Println(l.Print())
	head := l.Reserve()
	fmt.Println(l.Display(head))
	expect = 4
	if head.Data != expect {
		t.Errorf("expect:%d, actual:%d\n", expect, head.Data)
	}
	expect = 3
	if head.Next.Data != expect {
		t.Errorf("expect:%d, actual:%d\n", expect, head.Next.Data)
	}
}
func TestLinked_RemoveHead(t *testing.T) {
	l := NewLinked()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	fmt.Println(l.Print())
	So(t, 1, l.RemoveHead())
	So(t, 2, l.RemoveHead())
	So(t, 3, l.RemoveHead())
	So(t, 4, l.RemoveHead())
	fmt.Println(l.Print())
}
func TestLinked_RemoveTail(t *testing.T) {
	l := NewLinked()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	fmt.Println(l.Print())
	So(t, 1, l.RemoveTail())
	So(t, 2, l.RemoveTail())
	So(t, 3, l.RemoveTail())
	So(t, 4, l.RemoveTail())
	fmt.Println(l.Print())
}
func TestLinked_AddHead(t *testing.T) {
	var expect, actual interface{}
	l := NewLinked()
	l.Append(1)
	l.Append(2)
	expect = 1
	actual = l.Head.Data
	if expect != actual {
		t.Errorf("expect:%d, actual:%d\n", expect, actual)
	}
	expect = 2
	actual = l.Head.Next.Data
	if expect != actual {
		t.Errorf("expect:%d, actual:%d\n", expect, actual)
	}
	expect = nil
	actual = l.Head.Next.Next
	if l.Head.Next.Next != nil {
		So(t, expect, actual)
	}
}
func TestLinked_Remove(t *testing.T) {
	l := NewLinked()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	fmt.Println(l.Print())
	So(t, 2, l.Remove(2))
	fmt.Println(l.Print())
	So(t, 3, l.Remove(2))
	fmt.Println(l.Print())
	So(t, 4, l.Remove(2))
	fmt.Println(l.Print())
	So(t, nil, l.Remove(2))
	So(t, 1, l.Remove(1))
	fmt.Println(l.Print())
}
