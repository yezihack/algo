package main

//找出两个数之和
import (
	"fmt"
)

func main() {
	list := []int{5, 2, 33, 7, 11, 15}
	target := 9
	method1(list, target)
	method2(list, target)
	methed3(list, target)
}

//空间复杂度: O(n)
func methed3(list []int, target int) {
	fmt.Println("方法3:")
	total := 0
	i0, i1 := 0, 0
	m := make(map[int]int)
	for key, val := range list {
		total++
		//目标值 - 某一值 = 结果,判断结果是否在设置的m里,若存在则是
		diff := target - val
		if k, ok := m[diff]; ok {
			i0, i1 = k, key
		} else {
			m[val] = key
		}
	}
	fmt.Println("共循环次数", total)
	fmt.Printf("index0:%d, index1:%d\n", i0, i1)
}

//时间复杂度: O(n^2)
func method2(list []int, target int) {
	fmt.Println("方法二:")
	total := 0
	i0, i1 := 0, 0
	for i := 0; i < len(list); i++ {
		//j保存与i不重复比较,所以j = i + 1
		for j := i + 1; j < len(list); j++ {
			total++
			if target == (list[i] + list[j]) {
				i0, i1 = i, j
			}
		}
	}
	fmt.Println("共循环次数", total)
	fmt.Printf("index0:%d, index1:%d\n", i0, i1)
}

//双循环,时间复杂度为O(n^2)
func method1(list []int, target int) {
	fmt.Println("方法一:")
	total := 0
	i0, i1 := 0, 0
	for k, v := range list {
		for kk, vv := range list {
			total++
			if target == (v + vv) {
				i0, i1 = kk, k
			}
		}
	}
	fmt.Println("共循环次数", total)
	fmt.Printf("index0:%d, index1:%d\n", i0, i1)
}
