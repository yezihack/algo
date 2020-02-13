package _2_数学

import (
	"fmt"
	"math"
	"testing"
)

//https://leetcode-cn.com/problems/reverse-integer/
//第7道题
/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//思路, 解决反转问题, 一般采用栈方式解决, 栈的特点是先进后出.
func ReverseInteger(x int32) int32{
	var rev int32 = 0
	for x != 0 {
		mod := x % 10 //取余, 相当于弹出最高位.
		x = x / 10 //取余之后,需要将x除以10, 不断减小.
		fmt.Println("mod", mod, "x", x)
		//相当于入栈.这里有个风险,会出现元素的溢出.
		//如果rev大于最大32整数的话/10的话,才溢出.因为代码2处进行*10啦.所以这里要除以10.
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && mod > math.MaxInt32 % 10) { //防止溢出
			rev = 0
			break
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && mod < math.MaxInt32 % 10) {
			rev = 0
			break
		}
		rev = rev * 10 + mod //代码(2)
	}
	return rev
}
func TestReverseInteger(t *testing.T) {
	println(ReverseInteger(123))
	fmt.Println(math.MaxInt32, math.MaxInt32 % 10, math.MinInt32, math.MinInt32  % 10) //7, -8
}