package _3_动态规划

import (
	"fmt"
)

/*
题目： 给你一根长度为n的绳子，请把绳子剪成m段 (m和n都是整数，n>1并且m>1)每段绳子的长度记为k[0],k[1],…,k[m].
请问k[0] * k[1] …k[m]可能的最大乘积是多少？
例如，当绳子的长度为8时，我们把它剪成长度分别为2,3,3的三段，此时得到的最大乘积是18.
————————————————
版权声明：本文为CSDN博主「Jerry没有Tom」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/DEMON3344/article/details/85309025
*/

/*
动态规划一步一步的演算.
[0 1 2 3 0 0 0 0 0], i:4, j:1, i-j:3
[0 1 2 3 3 0 0 0 0], i:4, j:2, i-j:2
[0 1 2 3 4 0 0 0 0], i:5, j:1, i-j:4
[0 1 2 3 4 4 0 0 0], i:5, j:2, i-j:3
[0 1 2 3 4 6 0 0 0], i:6, j:1, i-j:5
[0 1 2 3 4 6 6 0 0], i:6, j:2, i-j:4
[0 1 2 3 4 6 8 0 0], i:6, j:3, i-j:3
[0 1 2 3 4 6 9 0 0], i:7, j:1, i-j:6
[0 1 2 3 4 6 9 9 0], i:7, j:2, i-j:5
[0 1 2 3 4 6 9 12 0], i:7, j:3, i-j:4
[0 1 2 3 4 6 9 12 0], i:8, j:1, i-j:7
[0 1 2 3 4 6 9 12 12], i:8, j:2, i-j:6
[0 1 2 3 4 6 9 12 18], i:8, j:3, i-j:5
[0 1 2 3 4 6 9 12 18], i:8, j:4, i-j:4
[0 1 2 3 4 6 9 12 18]
*/

func MaxCut(size int) int {
	if size < 2 {
		return 0
	}
	if size == 2 {
		return 1
	}
	if size == 3 {
		return 2
	}
	products := make([]int, size+1)
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3
	max := 0
	for i := 4; i <= size; i++ {
		max = 0
		for j := 1; j <= i/2; j++ {
			fmt.Printf("%v, i:%d, j:%d, i-j:%d\n", products, i, j, i-j)
			product := products[j] * products[i-j]
			if product > max {
				max = product
			}
			products[i] = max
		}
	}
	fmt.Println(products)
	max = products[size]
	return max
}

func TaiXinCut(n int) int {
	if n <= 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	ret := 1
	//尽量剪去3倍的绳子.
	for n > 3 {
		ret *= 3
		n -= 3
	}
	ret *= n
	return ret
}
func DpCut(n int) int {
	if n <= 1 {
		return 1
	}
	//申请切片, 存储已经计算好的值.
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	//从2开始计算.
	for i := 2; i <= n; i++ {
		max := 0
		for j := 1; j < i; j++ {
			a := j * (i - j)
			b := j * dp[i-j]
			fmt.Printf("i:%d, j:%d, max:%d, a:%d, dp:%d\n", i, j, max, a, b)
			max = max3(max, a, b)
		}
		dp[i] = max
	}
	fmt.Println(dp)
	return dp[n]
}
func max3(a, b, c int) int {
	ret := a
	if a < b {
		ret = b
	}
	if ret > c {
		return ret
	} else {
		return c
	}
}
