package Cache

import (
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	lru := NewLRUCache(3)
	//对GET判断
	node := lru.Get(1)
	if node != nil {
		t.Errorf("want: %v, errValue: %v\n", nil, node)
	}
	lru.Put(1, "lru")
	node = lru.Get(1)
	if node == nil {
		t.Errorf("want:not nil")
	}
	if node.Value != "lru" {
		t.Errorf("want: %s, errValue: %s", "lru", node.Value)
	}
	lru.Put(2, "cache")
	lru.Put(3, "golang")
	lru.Put(4, "php")
	node = lru.Get(1)
	if node.Value != nil {
		t.Errorf("want: %s, errValue: %s", "nil", node.Value)
	}
}
func TestLRUCache_Put(t *testing.T) {
	lru := NewLRUCache(5)
	for i := 1; i <= 5; i++ {
		lru.Put(i, i)
	}
	lru.Get(5)
	if node := lru.GetHead(); node.Key != 5 {
		t.Errorf("want: %d, errValue: %d", 5, node.Key)
	}
	lru.Get(2)
	if node := lru.GetHead(); node.Key != 2 {
		t.Errorf("want: %d, errValue: %d", 2, node.Key)
	}
	lru.Put(6, "sgfoot")
	lru.Get(6)
	if node := lru.GetHead(); node.Value != "sgfoot" {
		t.Errorf("want: %s, errValue: %s", "sgfoot", node.Value)
	}
	lru.Put(2, "github")
	lru.Get(2)
	if node := lru.GetHead(); node.Value != "github" {
		t.Errorf("want: %s, errValue: %s", "github", node.Value)
	}
	lru.Put(9, "99")
	node := lru.Get(100)
	if node != nil {
		t.Errorf("want: nil, err: %v", node)
	}
	lru.Print()
}

func TestLRUCache_Print(t *testing.T) {
	l := NewLRUCache(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	l.Print()
}
