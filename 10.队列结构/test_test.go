package queue_10

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestCycleQueue_Pop(t *testing.T) {
	q := NewCycleQ(4)
	src.Asset("a", t, true, q.Push(1))
	src.Asset("b", t, true, q.Push(2))
	src.Asset("c", t, true, q.Push(3))
	src.Asset("d", t, false, q.Push(4))
	src.Asset(1, t, 1, q.Pop())
	src.Asset(2, t, 2, q.Pop())
	src.Asset(3, t, 3, q.Pop())
	src.Asset(4, t, nil, q.Pop())
	src.Asset(5, t, 0, q.Length())
}
