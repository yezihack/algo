package dp

import (
	"fmt"
	"testing"
)

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//方法一: 动态规划
//爬1层,则是一种方法.爬2层,可以每次爬1层,也可以一次爬2层.
func ClimbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i ++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp)
	return dp[n-1]
}
func TestClimbStairs(t *testing.T) {
	tests := []struct{
		data int //数据
		expect int //预期值
	}{
		{2, 2},
		{3, 3},
	}
	for index, item := range tests {
		if actual := ClimbStairs(item.data); actual != item.expect {
			index ++
			t.Errorf("index:%d, expect:%d, actual:%d\n", index, item.expect, actual)
		}
	}
}