package linked

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestSingleLinked_InsertToHead(t *testing.T) {
	l := NewSingleLinked()
	for i := 0; i < 2; i ++ {
		l.InsertToHead(i+1)
	}
	l.Print()
	PrintLinked(Reverse(l.head))
}
func TestSingleLinked_InsertToTail(t *testing.T) {
	l := NewSingleLinked()
	for i := 0; i < 10; i ++ {
		l.InsertToTail(i+1)
	}
	l.Print()
}
// check linked is have cycle
func TestHasCycle(t *testing.T) {
	n1 := NewSingleNode(1)
	n2 := NewSingleNode(2)
	n3 := NewSingleNode(3)
	n4 := NewSingleNode(4)
	n1.next = n2
	n2.next = n3
	n3.next = n4
	src.Asset(1, t, false, HasCycle(n1))
	n4.next = n1
	src.Asset(2, t, true, HasCycle(n1))
}