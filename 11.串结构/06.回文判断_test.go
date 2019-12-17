package string_string

import (
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestPalindrome(t *testing.T) {
	src.Asset(1, t, true, Palindrome("aba"))
	src.Asset(2, t, true, Palindrome("abxba"))
	src.Asset(3, t, false, Palindrome("abx1ba"))
	src.Asset(4, t, false, Palindrome("ab2xba"))
}

func TestPalindromeByMid(t *testing.T) {
	src.Asset(1, t, true, PalindromeByMid("aba"))
	src.Asset(2, t, false, PalindromeByMid("abax"))
	src.Asset(3, t, true, PalindromeByMid("ababa"))
	src.Asset(4, t, true, PalindromeByMid("abaaba"))
	src.Asset(5, t, true, PalindromeByMid("1234554321"))
	src.Asset(6, t, true, PalindromeByMid("12345o54321"))
	src.Asset(7, t, false, PalindromeByMid("12345o154321"))
}