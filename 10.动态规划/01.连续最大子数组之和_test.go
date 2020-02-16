/*
 * @Author: 百里
 * @Date: 2020-02-12 17:16:16
 * @LastEditTime: 2020-02-12 17:17:10
 * @LastEditors: 百里
 * @Description:
 * @FilePath: \leetcode\10.动态规划\01.连续最大子数组之和.go
 * @https://github.com/yezihack
 */
package dp

import (
	"fmt"
	"testing"
)

//题目: 一维数组里有负数又有整数,求子数组连续求和最大.
//如[1, 2, -4, 2, 3], 连续求和最大是[2,3],和: 2+3=5

//方法一: 暴力法. 从0位置到尾都查看一遍,每计算一步都进行之前累加的数与当前数对比.并求出最大之和
func ContinueChildArraySum(arr []int) int {
	total := 0//求得最终的和
	for i := 0; i < len(arr); i ++ {
		sum := 0 //当前连续之和
		for j := i; j < len(arr); j ++ { //从i到n之间寻找最大值.
			sum += arr[j] //累加
			fmt.Printf("sum:%d, total:%d, cur:%d\n", sum, total, arr[j])
			//比较三者之间大小, 取最大值.
			total = max3(total, sum, arr[j]) //三个元素中查找最大值.赋值给total变量.
		}
	}
	return total
}
func TestContinueChildArraySum(t *testing.T) {
	if ret := ContinueChildArraySum([]int{1, -5, 2, -4, 2, 3}); ret != 5 {
		t.Errorf("actual:%d, expect:%d\n", ret, 5)
	}
	if ret := ContinueChildArraySum([]int{1, 1,-2, 3, 1}); ret != 4 {
		t.Errorf("actual:%d, expect:%d\n", ret, 4)
	}
	if ret := ContinueChildArraySum([]int{8, -8, -2, -2, -1, -1}); ret != 8 {
		t.Errorf("actual:%d, expect:%d\n", ret, 8)
	}
}

//方法二: 动态规划
//设置一个变量sum存储每一步的结果,如果当前值大于sum,则交换.
//再定义一个dp数组,存储当前最大的元素.
//时间O(n), 空间O(n)
func ContinueChildArraySumDp(arr []int) int {
	dp := make([]int, len(arr)) //动态规划数组
	sum := 0 //当前步骤之和
	for i := 0; i < len(arr); i ++ {
		sum += arr[i] //累加操作.
		if sum < arr[i] { //当前大于sum,才交换
			sum = arr[i]
		}
		if sum > dp[i] { //如果大于DP值,则存储起来.
			dp[i] = sum
		}
	}
	fmt.Println(dp)
	max := 0 //查找dp数组,取出最大值.
	for i := 0; i < len(dp); i ++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
func TestContinueChildArraySumDp(t *testing.T) {
	if ret := ContinueChildArraySumDp([]int{1, -5, 2, -4, 2, 3}); ret != 5 {
		t.Errorf("actual:%d, expect:%d\n", ret, 5)
	}
	if ret := ContinueChildArraySumDp([]int{1, 1,-2, 3, 1}); ret != 4 {
		t.Errorf("actual:%d, expect:%d\n", ret, 4)
	}
	if ret := ContinueChildArraySumDp([]int{8, -8, -2, -2, -1, -1}); ret != 8 {
		t.Errorf("actual:%d, expect:%d\n", ret, 8)
	}
}
//方法三: 动态规划
//题目说, 有正有负, 连续之和必定大于0. 如果使用一个变量累加时,与当前元素还小则进行交换.
//再重头计算.将每一次计算的结果存储在一个额外地变量里.最终得到最优解.
func ContinueChildArraySumDpV2(arr []int) int {
	var total, sum = 0, 0 //定义一个total为最优解, sum为当前连续最优解的变量.
	for i := 0; i < len(arr); i ++ {
		sum += arr[i] //求连续最优解之和
		fmt.Printf("cur:%d, sum:%d, total:%d\n", arr[i], sum, total)
		if arr[i] > sum { //如果当前值大于最优解则交换
			sum = arr[i] //交换
		}
		if sum > total { //如果最优解大于最终最优解则交换 .
			total = sum
		}
	}
	return total
}

func TestContinueChildArraySumDpV2(t *testing.T) {
	if ret := ContinueChildArraySumDpV2([]int{1, -2, 3, 1}); ret != 4 {
		t.Errorf("actual:%d, expect:%d\n", ret, 4)
	}
	return
	if ret := ContinueChildArraySumDpV2([]int{1, -5, 2, -4, 2, 3}); ret != 5 {
		t.Errorf("actual:%d, expect:%d\n", ret, 5)
	}
	if ret := ContinueChildArraySumDpV2([]int{1, 1,-2, 3, 1}); ret != 4 {
		t.Errorf("actual:%d, expect:%d\n", ret, 4)
	}
	if ret := ContinueChildArraySumDpV2([]int{8, -8, -2, -2, -1, -1}); ret != 8 {
		t.Errorf("actual:%d, expect:%d\n", ret, 8)
	}
}
func max3(a, b, c int) int {
	max := a
	if b > a {
		max = b
	}
	if c > max {
		return c
	}
	return max
}