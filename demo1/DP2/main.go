package main

import (
	"fmt"
	"sort"
)

func main() {
	best([]int{1, 10, 5, 2}, 16)
	best([]int{1, 10, 5, 2}, 106)
	best([]int{1, 10, 5, 2, 1}, 161)
	fmt.Println(10 % 2)
	fmt.Println(10 % 3)
	fmt.Println(10 % 5)
	fmt.Println(10 % 8)
}

func best(box []int, total int) {
	if total <= 0 {
		return
	}
	sort.Slice(box, func(i, j int) bool {
		return box[i] > box[j]
	})
	fmt.Println(box)
	newBox := make([]int, 0)
	for _, val := range box {
		if total < val {
			break
		}
		v1 := total / val
		v2 := total % val
		for i := 0; i < v1; i++ {
			newBox = append(newBox, val)
		}
		if v2 > 0 {
			total = v2
		} else {
			break
		}
	}
	//若还剩余则取最小值
	if total > 0 {
		newBox = append(newBox, box[len(box)-1])
	}
	MyPrint(newBox)
}

func MyPrint(b []int) {
	fmt.Println("---------------")
	total := 0
	fmt.Println(b)
	for _, val := range b {
		total += val
	}
	fmt.Println("共兑换:", total)
}
