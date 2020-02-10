package _3_动态规划

import (
	"fmt"
	"testing"
)

/**
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
 */

//暴力解法
func integerBreak(n int) int {
	if n < 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	ret := -1
	for i := 1; i < n; i ++ {
		ret = maxInt(ret, maxInt(i * (n - i), i * integerBreak(n-i)))
	}
	return ret
}
func TestIntegerBreak(t *testing.T) {
	fmt.Println(integerBreak(11))
}

//缓存中间过程
var cache map[int]int
func IntegerBreakCache(n int) int {
	if n < 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if _, ok := cache[n]; ok {
		return cache[n]
	}
	ret := -1
	for i := 1; i < n; i ++ {
		fmt.Printf("i:%d, a:%d, b:%d\n", i, i * (n-i), i * IntegerBreakCache(n-i))
		ret = maxInt(ret, maxInt(i * (n - i), i * IntegerBreakCache(n-i)))
	}
	cache[n] = ret
	return ret
}
func TestIntegerBreakCache(t *testing.T) {
	n := 10
	cache = make(map[int]int, n+1)
	fmt.Println(IntegerBreakCache(n))
}
//动态规划
func IntegerBreakDp(n int) int{
	if n == 2 {
		return 1
	}
	mem := make([]int, n+1) //申请一个存储中间件的数组.
	for i := 2; i <= n; i ++ {//从2开始,遍历到n的位置,包含n
		for j := 1; j < i; j ++ {//
			fmt.Printf("i:%d, j:%d, a:%d, b:%d, c:%d\n", i, j, mem[i], j * mem[i-j],j * (i-j))
			mem[i] = maxInt(mem[i], max(j * mem[i-j], j * (i-j)))
		}
	}
	return mem[n]
}
func TestIntegerBreakDp(t *testing.T) {
	fmt.Println(IntegerBreakDp(8))
}
func TestIntegerBreakTop(t *testing.T) {
	fmt.Println(IntegerBreakTop(8))
}
func IntegerBreakTop(n int) int {
	if n == 2 {
		return 1
	}
	dp := make([]int, n+1)

	for i := 3; i <= n; i ++ {
		for j := 1; j < i; j ++ {
			a := dp[i]
			b := j * dp[i-j]
			c := j * (i-j)
			dp[i] = maxInt(a, maxInt(b, c))
		}
	}
	return dp[n]
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMod(t *testing.T) {
	//求模
	a := 112
	fmt.Println(a % 2)//判断奇偶性, 使用%2
	fmt.Println(a & 1)//与1, 判断奇偶性. 如果是0则是偶数, 1是奇数.
}