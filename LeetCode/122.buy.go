package LeetCode

import "fmt"

//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

//贪心法.
func MaxProfit(prices []int) int {
	sum := 0
	for i := 0; i < len(prices); i ++ {
		if i + 1 < len(prices) && prices[i] < prices[i + 1] {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}
//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
//获取最大值与最小值. 相减.7,1,5,3,6,4
func MaxProfit2(prices []int) int {
	max := 0
	for i := 0; i < len(prices) -1; i ++ {
		for j := i + 1; j < len(prices); j ++ {
			diff := prices[j] - prices[i]
			if diff > max {
				max = diff
			}
		}
	}
	return max
}
func MaxProfit3(prices []int) int {
	min := 1 << 31 - 1
	max := 0
	for i := 0; i < len(prices); i ++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i] - min > max {
			max = prices[i] - min
		}
		fmt.Printf("min:%d, max:%d, price[i]:%d, p-min:%d\n", min,max, prices[i], prices[i]-min)
	}
	return max
}