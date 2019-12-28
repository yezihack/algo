package linked

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"testing"
)
func TestGetLinkedK(t *testing.T) {
	link := src.NewLinked()
	link.Append(1)
	link.Append(2)
	link.Append(3)
	link.Append(4)
	link.Append(5)
	link.Append(6)
	fmt.Println(link.Print())
	head := link.Head
	src.Asset(1, t, 5, GetLinkedK(head, 2).Data)
	src.Asset(2, t, 3, GetLinkedK(head, 4).Data)
	src.Asset(3, t, 6, GetLinkedK(head, 1).Data)
	src.Asset(4, t, 2, GetLinkedK(head, 5).Data)
	src.Asset(5, t, 1, GetLinkedK(head, 6).Data)
}

func TestGetLinkedKPoint(t *testing.T) {
	link := src.NewLinked()
	link.Append(1)
	link.Append(2)
	link.Append(3)
	link.Append(4)
	link.Append(5)
	link.Append(6)
	fmt.Println(link.Print())
	head := link.Head
	src.Asset(1, t, 5, GetLinkedKPoint(head, 2).Data)
	src.Asset(2, t, 3, GetLinkedKPoint(head, 4).Data)
	src.Asset(3, t, 6, GetLinkedKPoint(head, 1).Data)
	src.Asset(4, t, 2, GetLinkedKPoint(head, 5).Data)
	src.Asset(5, t, 1, GetLinkedKPoint(head, 6).Data)
}
func TestGetLinkedKCount(t *testing.T) {
	link := src.NewLinked()
	link.Append(1)
	link.Append(2)
	link.Append(3)
	link.Append(4)
	link.Append(5)
	link.Append(6)
	fmt.Println(link.Print())
	head := link.Head
	src.Asset(1, t, 5, GetLinkedKCount(head, 2).Data)
	src.Asset(2, t, 3, GetLinkedKCount(head, 4).Data)
	src.Asset(3, t, 6, GetLinkedKCount(head, 1).Data)
	src.Asset(4, t, 2, GetLinkedKCount(head, 5).Data)
	src.Asset(5, t, 1, GetLinkedKCount(head, 6).Data)
}