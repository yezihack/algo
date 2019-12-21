package linked

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)



func TestReverseLinked(t *testing.T) {
	l := src.NewLinked()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Print()
	head := l.Head
	result := ReverseLinked(head)
	t.Log(l.Display(result))
}

func TestReserveLinkedRecursion(t *testing.T) {
	l := NewSingleLinked()
	l.InsertToTail("a")
	l.InsertToTail("b")
	l.InsertToTail("c")
	l.InsertToTail("d")
	l.InsertToTail("e")
	l.Print()
	head := l.head
	result := ReserveLinkedRecursion(head)
	PrintLinked(result)
}
func TestInversion(t *testing.T) {
	l := src.NewLinked()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Print()
	head := l.Head
	result := Inversion(head)
	t.Log(l.Display(result))
}