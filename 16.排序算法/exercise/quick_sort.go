package exercise

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr) - 1)
}

func quickSort(arr []int, start, end int) {
	if start > end {
		return
	}
	index := partition(arr, start, end)
	quickSort(arr, start, index-1)
	quickSort(arr, index + 1, end)
}
func partition(arr []int, left, right int) int {
	pivot := arr[left]
	low, high := left, right
	for low < high  {
		for low < high && arr[high] > pivot {
			high --
		}
		arr[low] = arr[high]
		for low < high && arr[low] < pivot {
			low ++
		}
		arr[high] = arr[low]
	}
	arr[low] = pivot
	return low
}
