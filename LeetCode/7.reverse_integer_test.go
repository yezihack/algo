package LeetCode

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"math"
	"strconv"
	"testing"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
/*
示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
*/
func reverse(x int) int {
	is := true //标记是否为负数false, 正true
	if x < 0 {
		is = false
	}
	s := strconv.Itoa(x)
	t := ""
	//去掉负数
	if !is {
		s = s[1:]
	}
	//处理最末尾的0
	if s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	//倒转过来
	for i := len(s) - 1; i >= 0; i -- {
		t += string(s[i])
	}
	//字符串转换成整形

	toInt, err := strconv.ParseInt(t, 10, 32)
	if err != nil {
		return 0
	}
	if !is {
		return -int(toInt)
	}
	return int(toInt)
}
func TestReverse(t *testing.T) {
	src.Asset(1, t, 321, reverse(123))
	src.Asset(2, t, -321, reverse(-123))
	src.Asset(3, t, 21, reverse(120))
	src.Asset(4, t, 9646324351, reverse(1534236469))
}
//时间复杂度：O(\log(x))O(log(x))，xx 中大约有 \log_{10}(x)log
// (x) 位数字。
//空间复杂度：O(1)O(1)。
func reverseToInt(x int) int {
	y := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		fmt.Printf("pop:%d, x:%d, y:%d, max:%d, min:%d\n", pop, x, y, math.MaxInt32, math.MinInt32)
		if y > math.MaxInt32/10 || (y == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if y < math.MinInt32/10 || (y == math.MinInt32/10 && pop < -8) {
			return 0
		}
		y = y * 10 + pop
	}
	return y
}
func TestReverseToInt(t *testing.T) {
	src.Asset(1, t, 321, reverseToInt(123))
	//src.Asset(2, t, -321, reverseToInt(-123))
	//src.Asset(3, t, 21, reverseToInt(120))
	src.Asset(4, t, 2147483651, reverseToInt(1563847412))
	src.Asset(5, t, 9646324351, reverseToInt(1534236469))
}