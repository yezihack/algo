package dp

import (
	"math"
	"testing"
)

/*
斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
给定 N，计算 F(N)。

 

示例 1：

输入：2
输出：1
解释：F(2) = F(1) + F(0) = 1 + 0 = 1.
示例 2：

输入：3
输出：2
解释：F(3) = F(2) + F(1) = 1 + 1 = 2.
示例 3：

输入：4
输出：3
解释：F(4) = F(3) + F(2) = 2 + 1 = 3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fibonacci-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


//解法一: 递归
//推导出递推公式: 从上到下递推,如果想求f(n)的结果,则是f(n)=f(n-1)+f(n-2)
//time:2^N, space:O(n) 递归就是使用系统中的栈实现的.
func fib(n int) int{
	if n <= 1 {
		return n
	}
	return fib(n-2)+fib(n-1)
}
func TestFib(t *testing.T) {
	tests := []struct{
		data int //数据
		expect int //预期值
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}
	for index, item := range tests {
		if actual := fib(item.data); actual != item.expect {
			index ++
			t.Errorf("index:%d, expect:%d, actual:%d\n", index, item.expect, actual)
		}
	}
}
//解法二: 动态规划
//从下向上推导,最终求得结果
//Time:O(n), space:O(n)
func fibDp(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i ++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}
func TestFibDp(t *testing.T) {
	tests := []struct{
		data int //数据
		expect int //预期值
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}
	for index, item := range tests {
		if actual := fibDp(item.data); actual != item.expect {
			index ++
			t.Errorf("index:%d, expect:%d, actual:%d\n", index, item.expect, actual)
		}
	}
}

//直接使用推导公式
//f(n) =  f(n-2) + f(n-1) 即是前前项+前项=结果
//Time:O(n),space:O(1)
func fibMath(n int) int {

	if n <= 1 { //小于等于1则返回1
		return n
	}
	prepre, pre := 0, 1 //记录前前步, 前步的数字.
	for i := 2; i <= n; i ++ {
		prepre, pre = pre, pre + prepre //每次将前前项与前项相加,以次类推.求得结果.
	}
	return pre
}
func TestFibMath(t *testing.T) {
	tests := []struct{
		data int //数据
		expect int //预期值
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}
	for index, item := range tests {
		if actual := fibMath(item.data); actual != item.expect {
			index ++
			t.Errorf("index:%d, expect:%d, actual:%d\n", index, item.expect, actual)
		}
	}
}
//直接公式求解
//Time, space:O(1)
func FibMathV2(n int) int {
	var goldenRatio float64 = float64((1+math.Sqrt(5))/2)
	return int(math.Round(math.Pow(goldenRatio, float64(n))/math.Sqrt(5)))
	//var goldenRatio float64 = float64((1 + math.Sqrt(5)) / 2);
	//return int(math.Round(math.Pow(goldenRatio, float64(N)) / math.Sqrt(5)));
}
func TestFibMathV2(t *testing.T) {
	tests := []struct{
		data int //数据
		expect int //预期值
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}
	for index, item := range tests {
		if actual := FibMathV2(item.data); actual != item.expect {
			index ++
			t.Errorf("index:%d, expect:%d, actual:%d\n", index, item.expect, actual)
		}
	}
}