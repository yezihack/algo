/*
### 递归与尾递归的区别
> 递归调用时需要借用栈帧存储信息,栈帧由5个区域组成：
<br/>1.输入参数
<br/>2.返回值空间
<br/>3.计算表达式时用到的临时存储空间
<br/>4.函数调用时保存的状态信息
<br/>5.输出参数
1. 递归:递归前进时保存大量的栈帧信息,只有返回时才会释放
1. 尾递归:当递归调用是整个函数体中最后执行的语句且它的返回值不属于表达式的一部分时，这个递归调用就是尾递归
尾递归函数的特点是在回归过程中不用做任何操作
*/
package Recurive

import (
	"fmt"
	"testing"
)

//非尾递归
func Fact(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 || n == 0 {
		return 1
	}
	return n * Fact(n-1) //n变量会创建一个栈帧保存起来,直到返回栈时才会释放,这样消费很多空间存储
}

func TestFact2(t *testing.T) {
	fmt.Println(Fact(10))
}

//尾递归,编译器对其做了优化,效率可以与循环相媲美
func FactTail(n, result int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return result
	}
	return FactTail(n-1, n*result) //这是没有使用变量存储,也就不会创建变量栈帧,而且将结果当参数覆盖每一次计算,最后返回结果
}
//斐波那契数列
//1, 1, 2, 3, 5, 8, 13, 21
func FactV2(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FactV2(n - 1) + FactV2(n - 2)
}
func TestFactV2(t *testing.T) {
	fmt.Println(FactV2(2))
	fmt.Println(FactV2(3))
	fmt.Println(FactV2(4))
	fmt.Println(FactV2(5))
	fmt.Println(FactV2(6))
}

func FabV3(n int) int {
	if n < 0 {
		return 0
	}
	x, y := 0, 1
	for i := 0; i < n; i ++ {
		x, y = y, x + y
	}
	return x
}
func TestFactV3(t *testing.T) {
	fmt.Println(FabV3(1))
	fmt.Println(FabV3(2))
	fmt.Println(FabV3(3))
	fmt.Println(FabV3(4))
	fmt.Println(FabV3(5))
	fmt.Println(FabV3(6))
}
