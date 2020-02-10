package queue_10

import (
	"fmt"
	"testing"
)

/**
约瑟夫斯问题（有时也称为约瑟夫斯置换），是一个出现在计算机科学和数学中的问题。在计算机编程的算法中，类似问题又称为“约瑟夫环”，也有的地方叫做“丢手绢”。
有编号从1 到n的 n个人围坐成一圈。从编号为1的人开始报数，报到 m 的人出局，
下一位再从 1 开始报数，报到 m 的人出局，……如此持续，
直到剩下一人为止，假设此人的原始编号是x。给定 n和 m，求出x。
 */
func YueSeHuCycle(n, m int) int {
	if n <= 0 {
		return 0
	}
	if m <= 0 {
		return 0
	}
	cycle := make([]int, n) //初使一个圈子
	count := 0 //记录已经出圈的人
	number := 0 //报数器.
	//所有人入圈子
	for i := 0; i < n; i ++ {
		cycle[i] = 1
	}
	for count < n-1 { //留一个人
		for i := 0; i < n; i ++ { //从头数到尾
			if cycle[i] == 1 { //未出圈子的人
				number ++ //数数
				if number == m { //报数到m就出圈子
					count ++ //记录已经出圈子的数量
					number = 0 //又从头数.
					cycle[i] = 0//标记出圈的人
				}
				if count == n - 1 {
					break
				}
			}
		}
	}
	fmt.Println(cycle)
	//找出数组标记为1的人
	for i := 0; i < n; i ++ {
		if cycle[i] == 1 {
			fmt.Println("剩余最后一个人:", i+1)
			return i +1
		}
	}
	return 0
}
func TestYueSeHuCycle(t *testing.T) {
	fmt.Println(YueSeHuCycle(41, 2))
}

func JosCycle(n, m int) int {
	if n <= 0 || m <= 0 {
		return 0
	}
	//初使化
	cycle := make([]int, n)
	for i := 0; i < n; i ++ {
		cycle[i] = 1
	}
	var (
		count int //记录出局人数
		number int//记录报数次数
	)
	for count < n -1 {
		for i := 0; i < n; i ++ { //不断循环
			if cycle[i] == 1 { //必须是未出局的人
				number ++ //开始报数
				if number == m {
					number = 0//重新开始报报数
					count ++ //开始记录出局数量
					cycle[i] = 0//设置出局
				}
				if count == n - 1 {
					break
				}
			}
		}
	}
	fmt.Println("finish", cycle)
	for i := 0; i < n; i ++ {
		if cycle[i] == 1 {
			return i + 1
		}
	}
	return 0
}
func TestJosCycle(t *testing.T) {
	fmt.Println(JosCycle(5, 2))
}