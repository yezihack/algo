package Recurive

import (
	"fmt"
	"testing"
)

/*
小青蛙跳台阶, 可以跳一个台阶, 也可以跳二个台阶.总共有多少种跳法.
分析一下: 如果跳一个台阶, 还剩n-1个台阶没跳, 如果跳一次跳二个台阶, 还剩n-2个台阶没跳.
推导出递推公式f(n) = f(n-1) + f(n-2)
解题思路三步走写递归算法:
1.首先明白定义递归函数功能?计算小青蛙一共有多少跳法.
2.找出终止条件, 先分解最小跳法,1个台阶的走法是1种方式, 2个台阶的走法是2种方式.
3.将递推公式写出来.
*/
//求解台阶走法. 1个台阶的走法是1种方式, 2个台阶的走法是2种方式.
func Step(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return Step(n - 1) + Step(n -  2)
}

func StepV2(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return StepV2(n - 1) + StepV2(n - 2)
}
func TestStepV2(t *testing.T) {
	fmt.Println(StepV2(6))
}

var cache = make(map[int]int)
func StepMap(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if val, ok := cache[n]; ok {
		return val
	}
	ret := StepMap(n-1) + StepMap(n-2)
	cache[ret] = ret
	return ret
}
func StepF(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	ret := 0
	pre := 2
	prepre := 1
	for i := 3; i <= n; i ++ {
		ret = pre + prepre
		prepre = pre
		pre = ret
	}
	return ret
}

func StepNo(n int) int{
	if n < 0{
		return 0
	}
	if n < 3 {
		return n
	}
	result := 0
	prepre := 1
	pre := 2
	for i := 3; i < n;i ++ {
		result = pre + prepre
		prepre = pre
		pre = result
	}
	return result
}
func TestStepNo(t *testing.T) {
	fmt.Println(StepNo(6))
}