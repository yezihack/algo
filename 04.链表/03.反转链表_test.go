package linked

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestIsPalindrome1(t *testing.T) {
	l := NewSingleLinked()
	l.InsertToHead("a")
	l.InsertToHead("b")
	l.InsertToHead("c")
	l.InsertToHead("b")
	l.InsertToHead("a")
	l.Print()
	src.Asset(1, t, true, IsPalindrome1(l))
}
