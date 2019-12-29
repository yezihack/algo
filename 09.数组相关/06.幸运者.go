package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

/*
* 约瑟夫环问题：
 * N个人排成一队，从1到5轮流报数，报5的是幸运者，出列
 * 报到队尾后，从队尾接着报，依此循环
 * 问：排在队尾的人是第几名幸运者？
 * N为小于100000的正整数
 * 例如：1人排成一队，他就是第一名幸运者
 * 3人排成一队，队尾是第二名幸运者
 * 5人排成一队，队尾是第一名幸运者
 * 8人排成一队，队尾是第三名幸运者
 *
 * 输入：队伍的总人数————20
 * 输出：队尾者的幸运编号——————4
*/
func main() {
	if GoodLuckNumber(1, 5) != 1 {
		log.Error("Not luck number")
	}
	if GoodLuckNumber(3, 5) != 2 {
		log.Error("Not luck number")
	}
	if GoodLuckNumber(5, 5) != 1 {
		log.Error("Not luck number")
	}
	if GoodLuckNumber(8, 5) != 3 {
		log.Error("Not luck number")
	}
	if GoodLuckNumber(20, 5) != 4 {
		log.Error("Not luck number")
	}
	if JosephRingGoodLuck(1, 5) != 1 {
		log.Error("Not luck number")
	}
	if JosephRingGoodLuck(3, 5) != 2 {
		log.Error("Not luck number")
	}
	if JosephRingGoodLuck(5, 5) != 1 {
		log.Error("Not luck number")
	}
	if JosephRingGoodLuck(8, 5) != 3 {
		log.Error("Not luck number")
	}
	if JosephRingGoodLuck(20, 5) != 4 {
		log.Error("Not luck number")
	}
	fmt.Println("ok")
}
func GoodLuckNumber(n, k int) int {
	if n <= 0 || k <= 0 {
		return -1
	}
	if n == 1 {
		return 1
	}
	cycle := make([]int, n)
	for i := 0; i < n; i ++ {
		cycle[i] = i + 1
	}
	record := 0 //记录踢出多少个人,如果是队尾元素则是目标值.
	number := 0 //报数
	index := 0 //数组指针
	for record < n - 1 {
		if cycle[index] != 0 { //不为0则开始数数
			number ++
		}
		if number == k { //报数与k相等,则表示需要对数据进行标记啦
			cycle[index] = 0 //标记为0
			record ++ //记录标记为0的数量
			if index == n -1 { //如果是数组最后一个元素,则求得了目标值
				return record
			}
			number = 0 //重新开始数.
		}
		index ++ //移动数组指针
		if index == n { //如果到了数组尾部则从头开始.
			index = 0
		}
	}
	return -1
}
func JosephRingGoodLuck(n, k int) int {
	//边界处理
	if n <= 0 || k <= 0 {
		return -1
	}
	if n == 1 {
		return 1
	}
	//申请一个环
	ring := make([]bool, n)
	for i := 0; i < n; i ++ {
		ring[i] = true
	}
	index := 0 //环的指针
	number := 0 //报数: 1,2,3,k
	level := 0 //排名
	for {
		if ring[index] { //环元素为true时
			number ++ //开始报数
			if number == k { //报数的值等于k
				level ++ //排名+=1
				number = 0 //重新开始
				ring[index] = false //标记环中元素,下次不参与报数
				if index == n - 1 { //如果是数组最末尾则返回排名level
					return level
				}
			}
		}
		index ++ //移动指针
		if index == n { //如果移动到尾部则从头再开始.
			index = 0
		}
	}
	return -1
}
