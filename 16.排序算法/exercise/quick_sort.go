package exercise

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr) - 1)
}

func quickSort(arr []int, start, end int) {
	if start < end {
		index := partition(arr, start, end) //交换数据,并返回基准位置
		quickSort(arr, start, index-1) //对左子表递归排序
		quickSort(arr, index + 1, end) //对右子表递归排序
	}
}
//以中间值为准,进行交换数据
func partition(arr []int, left, right int) int {
	pivot := arr[left] //设置一个基准值
	low, high := left, right //申请一个低位指针,一个高位指针
	for low < high  { //必须低位小于高位.循环才成立
		for low < high && arr[high] > pivot { //如果基准值小于高位值,则高位指针向左移动
			high -- //高位指针向左移动
		}
		arr[low] = arr[high] //小于基准值的移动到最左边去
		for low < high && arr[low] < pivot {//如果基准值大小低位值,则低位指针向右移动
			low ++ //低位指针向右移动
		}
		arr[high] = arr[low]//大于基准值的移动到最右边去
	}
	arr[low] = pivot//将基准值置换到中间位置
	return low//返回中间位置
}
