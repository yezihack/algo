package Cache

import (
	"fmt"
	"testing"
)

func TestLFUCache_Put(t *testing.T) {
	lfu := NewLFUCache(3)
	lfu.Put(1, 1)
	lfu.Print()
	lfu.Put(2, 2)
	lfu.Print()
	lfu.Get(2)
	lfu.Get(2)
	lfu.Get(2)
	lfu.Get(2)
	lfu.Print()
	lfu.Put(3, 3)
	lfu.Print()
	lfu.Put(4, 4)
	lfu.Print()
	lfu.Get(1)
	lfu.Get(1)
	lfu.Get(1)
	fmt.Println(lfu.Get(1))
	lfu.Print()
}
func TestLFUCache_Get(t *testing.T) {
	l := NewLFUCache(10)
	for i := 0; i < 10; i++ {
		l.Put(i, i)
	}
	l.Get(3)
	l.Get(3)
	if l.MaxFreq() != 3 {
		t.Errorf("want: %v, err:%v\n", 3, l.MaxFreq())
	}
	l.Put(10, 10)
	l.Get(10)
	l.Get(10)
	node := l.Get(10)
	if node.Value != 10 {
		t.Errorf("want:%v, err:%v\n", 10, node.Value)
	}
	if l.MaxFreq() != 4 {
		t.Errorf("want: %v, err:%v\n", 4, l.MaxFreq())
	}
	l.Print()
}
