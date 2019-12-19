package _6_排序算法

import "fmt"

//快速排序
func FastSort(arr []int, left, right int) {
	if left < right {
		//找到中心轴排好序的位置
		pivot := partition(arr, left, right)
		//对低子表递归排序
		FastSort(arr, left, pivot - 1)
		//对高子表递归排序
		FastSort(arr, pivot + 1, right)
	}
}
//对数据元素进行左右调整
//小于中心值则向左边调换
//大于中心值则向右边调换
//借用low, high指针,一个左,一个右.
func partition(arr []int, low, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high && arr[high] > pivot {
			high --
		}
		arr[low] = arr[high]
		for low < high && arr[low] < pivot {
			low ++
		}
		arr[high] = arr[low]
	}
	//跳出for表示low,high指针重合.
	arr[low] = pivot
	fmt.Printf("pivot:%d,low:%d,high:%d,%v\n",pivot, low,high,arr)
	return low
}
/*************快速排序另一个版本.一个函数搞定*****************/
func QuickSort(arr []int, left, right int) {
	pivot := arr[left]//找个中心点
	low, high := left, right //使用二个指针,一个在左,一个在右
	for low < high {//不重合的状态下,循环
		for low < high && arr[high] >= pivot {//右边一直大于中心值,则向前移动high指针
			high --
		}
		arr[low] = arr[high]//当小于中心点,则让数据交互到最低点
		for low < high && arr[low] <= pivot {//左边的一直小于中心值,则向前移动low指针
			low ++
		}
		arr[high] = arr[low]//当大于中心点,则让数据交互到最高点
	}
	arr[low] = pivot//将中心点的值填入到中心位置
	fmt.Printf("low:%d, high:%d, left:%d, right:%d, low-left:%d, right-low:%d, %v\n",
						low, high, left, right, low-left, right-low, arr)
	if low - left > 1 {
		QuickSort(arr, left, low -1)
	}
	if right - low > 1 {
		QuickSort(arr, low+1, right)
	}
}

/*****************3*******************/
func QuickSort3(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := _partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func _partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index ++
		}
	}
	index --
	fmt.Printf("left:%d,right:%d,pivod:%d,index:%d,%v\n", left, right, pivot, index, arr)
	swap(arr, pivot, index)
	return index
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


