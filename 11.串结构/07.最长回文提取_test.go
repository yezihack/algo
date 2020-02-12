package string_string

import (
	"fmt"
	"strings"
	"testing"
)

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	dp := make(map[string]struct{}, 0)
	for i := 1; i <= len(s); i ++ {
		for j := 0; j < i; j ++ {
			if IsPalindrome(s[j:i]) {
				if _, ok := dp[s[j:i]]; !ok {
					dp[s[j:i]] = struct{}{}
				}
			}
		}
	}
	fmt.Println(dp)
	maxLongStr := ""
	for s := range dp {
		if len(s) > len(maxLongStr) {
			maxLongStr = s
		}
	}
	return maxLongStr
}
func TestLongestPalindrome(t *testing.T) {
	fmt.Println(LongestPalindrome("ac"))
	fmt.Println(LongestPalindrome("bb"))
	fmt.Println(LongestPalindrome("babad"))
	str := strings.Repeat("a", 1000)
	fmt.Println(LongestPalindrome(str))
}
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	//向中间逼近,判断
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left ++
		right --
	}
	return true
}
func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("abba") {
		t.Error("not palindrome")
	}
	if IsPalindrome("ababcc") != false {
		t.Error("not palindrome")
	}
}

func LongPalindrome(s string) string {
	size := len(s)
	if size < 2 {
		return s
	}
	dp := make([][]bool, size)
	for i := 0; i < size; i ++ {
		dp[i] = make([]bool, size)
		dp[i][i] = true
	}
	maxLen := 1
	start := 0
	for j := 1; j < size; j ++ {
		for i := 0; i < j; i ++ {
			if s[i] == s[j] {
				if j - i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] {
				curLen := j - i + 1
				if curLen > maxLen {
					maxLen = curLen
					start = i
				}
			}
		}
	}
	for i := 0; i < len(dp); i ++ {
		for j := 0; j < len(dp); j ++ {
			fmt.Printf("i:%d,j:%d %v ", i,j,dp[i][j])
		}
		fmt.Println()
	}
	fmt.Println("start", start, start+maxLen)
	return s[start:start+maxLen]
}
func TestLongPalindrome(t *testing.T) {
	fmt.Println(LongPalindrome("babad"))
}