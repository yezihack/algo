/*
 * @Author: 百里
 * @Date: 2020-02-11 08:46:09
 * @LastEditTime : 2020-02-12 17:37:08
 * @LastEditors  : 百里
 * @Description:
 * @FilePath: \leetcode\20.剑指offer\15.二进制1的个数_test.go
 * @https://github.com/yezihack
 */
package 剑指offer

import (
	"fmt"
	"testing"
)
//求一个数字中的二进制里的1的个数.
//比如10的数字二进制是1010, 含有二个1,则返回2
//思路: 设置一个flag=1,即二进制是0001与1010进行与运算
/*
与运算即: 0&1=0, 1&0=0, 0&0=0, 1&1=0 只有二个1,1才是1,顺便复习一下, 或, 非吧. 易记法:双1才是1,其它都为0
或运算即: 0|1=1, 1|0=1, 0|0=0, 1|1=1 只要有个1就是1,易记法: 有1就是1,其它都为0
非运算即: ^1=0, ^0=1
异或运算即: 0^1=0, 1^0=0, 0^0=1, 1^1=1 易记法:相同才为1, 其它为0

1010  	1010	1010	1010
0001	0010	0100	1000
-----------------------------进行与&运算
0000	0010	0000	1000
 */

//从最底位开始进行, n和flag进行位移操作. 然后向左移.
func NumberOf1InBinaryV1(n int) int {
	count := 0
	flag := 1
	for flag != 0 {
		if (n & flag) != 0 {
			count++
		}
		fmt.Println(flag)
		flag = flag << 1
	}
	fmt.Println(flag)
	return count
}

func TestNumberOf1InBinaryV1(t *testing.T) {
	t.Log("result")
	if NumberOf1InBinaryV1(10) != 2 {
		t.Error("result not equal 2")
	}
}
func NumberOf1InBinaryV2(n int) int {
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

func TestNumberOf1InBinaryV2(t *testing.T) {
	t.Log("result")
	if NumberOf1InBinaryV2(9) != 2 {
		t.Error("result not equal 2")
	}

}
func Find1OfInBinary(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n <<= 1
	}
	return count
}
func TestFind1OfInBinary(t *testing.T) {
	fmt.Println(Find1OfInBinary(10))
}
