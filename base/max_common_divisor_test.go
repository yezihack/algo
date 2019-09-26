package base

import (
	"fmt"
	"testing"
)

//最大公约数
func recursionCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}
	return recursionCommonDivisor(b, a%b)
}
func forCommonDivisor(a, b int) int {
	fmt.Printf("first:a:%d, b:%d \n", a, b)
	for b != 0 {
		ret := a % b
		a = b
		b = ret
		fmt.Printf("a:%d, b:%d，ret: %d\n", a, b, ret)
	}
	return a
}
func TestRecursionCommonDivisor(t *testing.T) {
	var a, b int
	a, b = 105, 24
	fmt.Println(recursionCommonDivisor(a, b))
	fmt.Println(forCommonDivisor(a, b))
	fmt.Println(forCommonDivisor(1680, 640))
	fmt.Println(105 % 24)
	fmt.Println(24 % 105)

}
