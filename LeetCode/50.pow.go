package LeetCode

import (
	"fmt"
	"math"
)

//实现 pow(x, n) ，即计算 x 的 n 次幂函数。
func pow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
func Pow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	return x * pow(x, n-1)
}

func Pow3(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	//处理n为负的情况.负次方就是数倒数.
	if n < 0 {
		x = 1/x
		n = -n
	}
	var result float64
	result = x
	for i := n; i > 1; i -- {
		result = result * x
	}
	return result
}

func Pow4(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1/x
		n = -1
	}
	var result float64
	result = 1
	i := n
	for i != 0 {
		if i % 2 != 0 {
			result *= x
		}
		x *= x
		i = i/2
		fmt.Printf("i:%d, x:%0.3f, result:%0.3f\n", i, x, result)
	}
	return result
}
//分治法. 计算一半即可
func Pow5(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1/x
		n = -n
	}
	half := Pow5(x, n/2)
	if n & 1 == 1 {
		return half * half * x
	} else {
		return half * half
	}
}
