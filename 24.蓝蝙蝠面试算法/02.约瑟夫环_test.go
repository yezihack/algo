package _4_蓝蝙蝠面试算法

import (
	"fmt"
	"testing"
)

/*
约瑟夫环（约瑟夫问题）是一个数学的应用问题：
已知 n 个人（以编号1，2，3…n分别表示）围坐在一张圆桌周围。
从编号为 k 的人开始报数，数到 m 的那个人出圈；
他的下一个人又从 1 开始报数，数到 m 的那个人又出圈；依此规律重复下去，直到剩余最后一个胜利者。
*/

func JosCycle(n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}
	//定义一个数组,模拟环
	cycle := make([]int, n)
	//初使环, 未踢出的人标记为1, 踢出的标记为0
	for i := 0; i < n; i ++ {
		cycle[i] = 1
	}
	var (
		count = 0 //定义一个用于踢出人的数量
		number = 0//定义一个报数的变量.
	)
	for count < n -1 { //如果踢出的人
		for i := 0; i < n; i ++ {
			if cycle[i] == 1 {
				number ++
				if number % m == 0 {
					number = 0
					cycle[i] = 0
					count ++
				}
				if count == n-1 {
					break
				}
			}
		}
	}
	fmt.Println(cycle)
	for i := 0; i < n;i ++ {
		if cycle[i] == 1 {
			return i + 1
		}
	}
	return 0
}
func TestJosCycle(t *testing.T) {
	fmt.Println(JosCycle(4, 2))
}