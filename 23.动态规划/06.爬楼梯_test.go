package _3_动态规划

import (
	"fmt"
	"testing"
)

//n阶台队的楼梯, 可以一次1阶或2阶, 求最多有多少种爬法
//先求状态的定义. f(n), 走到最后一个台阶的走法. 就是f(n)表示.
// f(n)倒数第一个台阶可以到达, 也可以是倒数第二个台阶到达.
// 所以推荐公式: f(n) = f(n-1) + f(n-2)
//公式:f(n) = f(n-1) + f(n-2).

func Step(n int) int {
	if n < 1 {
		return 0
	}
	if n < 2 {
		return 1
	}
	arr := make([]int, n)
	arr[0] = 1 //第一个台阶的走法,只有1种
	arr[1] = 2 //第二个台阶的走法,有二种: 1,1; 2
	for i := 2; i < n; i ++ { //从第3个台阶开始走.
		arr[i] = arr[i-1] + arr[i-2] //根据递推公式:f(n)=f(n-1) + f(n-2)
	}
	fmt.Println(arr)
	return arr[n-1] //得到最后一个元素
}
func TestStep(t *testing.T) {
	fmt.Println(Step(3))
	fmt.Println(Step(6))
}
func ClimbStairs(n int) int{
	if n < 0 {
		return 0
	}
	x, y := 1, 1
	for i := 1; i < n; i ++ {
		x, y = y, x + y
		fmt.Printf("x:%d, y:%d\n", x, y)
	}
	return y
}
func TestClimbStairs(t *testing.T) {
	fmt.Println(ClimbStairs(3))
	fmt.Println(ClimbStairs(6))
}