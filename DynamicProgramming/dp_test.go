package DynamicProgramming

import (
	"testing"
)

//算法
// 先正常思维
// 逆向思维
// 递推思维 逐一推导结果,在此发现规律.

//动态规划算法Dynamic Programming(递推) 动态递推
//对问题求解时,将问题分解成若干个子问题,
/*
求解方法:
1.递归 + 记忆化 -> 递推
2.状态的定义
3.状态的转移方程
4.最优子结构
*/

//第一个题 ,非波纳契
//第一种解法:递归
func fab(n int) int {
	if n <= 2 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}
func TestFab(t *testing.T) {
	var r int
	r = fab(5)
	if r != 5 {
		t.Errorf("result is err, errResult: %d", r)
	}
	r = fab(6)
	if r != 8 {
		t.Errorf("result is err, errResult: %d", r)
	}
	r = fab(8)
	if r != 21 {
		t.Errorf("result is err, errResult: %d", r)
	}
}

//第二种解法: 递归+记忆化
//因为有重复的计算, 我们将其放入缓存中
func fabCache(n int, cache map[int]int) int {
	if n <= 2 {
		return 1
	}
	if _, ok := cache[n]; !ok {
		cache[n] = fabCache(n-1, cache) + fabCache(n-2, cache)
	}
	return cache[n]
}
func TestFabCache(t *testing.T) {
	var r int
	cache := make(map[int]int)
	r = fabCache(5, cache)
	if r != 5 {
		t.Errorf("result is err, errResult: %d", r)
	}
	r = fabCache(6, cache)
	if r != 8 {
		t.Errorf("result is err, errResult: %d", r)
	}
	r = fabCache(8, cache)
	if r != 21 {
		t.Errorf("result is err, errResult: %d", r)
	}
}

//第三种解法: 递推法
//首先找出规律公式 f(n) = f(n-1) + f(n-2)
func fabProgramming(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func TestFabProgramming(t *testing.T) {
	var r int
	r = fabProgramming(5)
	if r != 5 {
		t.Errorf("result is err, errResult: %d", r)
	}
	r = fabProgramming(6)
	if r != 8 {
		t.Errorf("result is err, errResult: %d", r)
	}
	r = fabProgramming(8)
	if r != 21 {
		t.Errorf("result is err, errResult: %d", r)
	}
}
func fabProgramming2(n int) int {
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}
func TestFabProgramming2(t *testing.T) {
	var r int
	r = fabProgramming2(5)
	if r != 5 {
		t.Errorf("result is err, errResult: %d", r)
	}
	r = fabProgramming2(6)
	if r != 8 {
		t.Errorf("result is err, errResult: %d", r)
	}
	r = fabProgramming2(8)
	if r != 21 {
		t.Errorf("result is err, errResult: %d", r)
	}
}
