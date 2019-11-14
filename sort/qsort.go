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

func FastSort(arr []int) {
	_fastSort(arr, 0, len(arr)-1)
}
func _fastSort(arr []int, left, right int) {
	if left < right {
		index := _partition(arr, left, right)
		_fastSort(arr, left, index - 1)
		_fastSort(arr, index + 1, right)
	}
}
func _partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i ++ {
		if arr[i] < arr[pivot] {
			_swap(arr, i, index)
			index ++
		}
	}
	index --
	_swap(arr, pivot, index)
	return index
}
func _swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}