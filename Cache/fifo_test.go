package Cache

import (
	"fmt"
	"testing"
)

func TestFIFO_Put(t *testing.T) {
	fifo := NewFIFO(3)
	fifo.Put(1, 1)
	fifo.Print()
	fifo.Put(2, 2)
	fifo.Print()
	fifo.Put(3, 3)
	fifo.Print()
	fifo.Put(4, 4)
	fifo.Print()
	fifo.Put(5, 5)
	fifo.Print()
	fifo.Put(3, "我的一个3")
	fifo.Print()
}
func TestFIFO_GetSize(t *testing.T) {
	f := NewFIFO(2)
	f.Put(0, 0)
	if f.GetSize() != 1 {
		t.Errorf("size:%d, errSize: %d\n", 1, f.GetSize())
	}
	f.Put(1, 1)
	if f.GetSize() != 2 {
		t.Errorf("size:%d, errSize: %d\n", 2, f.GetSize())
	}
	f.Put(3, 3)
	f.Print()
	if f.GetSize() != 2 {
		t.Errorf("size:%d, errSize: %d\n", 2, f.GetSize())
	}
}
func TestFIFOCache_Get(t *testing.T) {
	f := NewFIFO(10)
	f.Put(1, "abc")
	node := f.Get(1)
	if node.Key != 1 {
		t.Errorf("want: %d, err: %v", 1, node.Key)
	}
	if f.GetSize() != 1 {
		t.Errorf("want: %d, err: %v", 1, f.GetSize())
	}
	for i := 0; i < 100; i++ {
		f.Put(i, fmt.Sprintf("i:%d", i))
	}
	if f.GetSize() != 10 {
		t.Errorf("want: %d, err:%v", 10, f.GetSize())
	}
	fmt.Println(f.Get(99))
	fmt.Println(f.Get(1001))
}
