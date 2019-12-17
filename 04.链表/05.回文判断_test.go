package linked

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)

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
