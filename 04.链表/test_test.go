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
	Print(Reverse(l.head))
}
func TestSingleLinked_InsertToTail(t *testing.T) {
	l := NewSingleLinked()
	for i := 0; i < 10; i ++ {
		l.InsertToTail(i+1)
	}
	l.Print()
}
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
func TestCheckPalindrome(t *testing.T) {
	src.Asset(1, t, true, CheckPalindrome("abccba"))
	src.Asset(2, t, true, CheckPalindrome("aba"))
	src.Asset(3, t, false, CheckPalindrome("aabb"))
	src.Asset(4, t, true, CheckPalindrome("aabaa"))
}
func TestCheckPalindrome2(t *testing.T) {
	src.Asset(1, t, true, CheckPalindrome2("aba"))
	src.Asset(2, t, true, CheckPalindrome2("aabaa"))
	src.Asset(3, t, true, CheckPalindrome2("ccppcc"))
	src.Asset(4, t, false, CheckPalindrome2("abc"))
	src.Asset(5, t, false, CheckPalindrome2(""))
	src.Asset(6, t, true, CheckPalindrome2("程序序程"))
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