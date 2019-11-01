package sort

import "fmt"

func QSort(list []int, left, right int) {
	pivot := list[left]
	i, j := left, right
	for i < j {
		for i < j && list[j] > pivot {
			j--
		}
		if i < j {
			list[i] = list[j]
			i++
		}
		for i < j && list[i] < pivot {
			i++
		}
		if i < j {
			list[j] = list[i]
			j--
		}
	}
	list[i] = pivot
	fmt.Printf("left:%d, right:%d, i:%d, j :%d\n", left, right, i, j)
	//QSort(list, left, i-1)
	//QSort(list, i+1, right)
	if i-left > 1 {
		QSort(list, left, i-1)
	}
	if right-i > 1 {
		QSort(list, i+1, right)
	}
}
