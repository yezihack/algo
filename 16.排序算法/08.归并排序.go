package _6_排序算法

//归并是采用分治思想.先折半数组,然后再折半.直到一个元素的数组.然后进行合并一个数组
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr)/2 //一分为二
	left := MergeSort(arr[0:mid]) //将数组的前半部分递归
	right := MergeSort(arr[mid:]) //将数组的后半部分递归
	result := merge(left, right)  //分解的传递给合并函数
	return result
}
//将分解的小数组合并成一个大数组
func merge(left, right []int) []int {
	result := make([]int, 0)//申请一个临时空间
	m, n := 0, 0 //申请二个指针
	l, r := len(left), len(right) //申请二个数组的边界指针
	for m < l && n < r {
		if left[m] < right[n] { //小于则加入临时数组
			result = append(result, left[m])
			m ++
		} else {
			result = append(result, right[n])
			n ++
		}
	}
	if m < l {//将未移动完的数据直接加入临时数组
		result = append(result, left[m:]...)
	}
	if n < r {//将未移动完的数据直接加入临时数组
		result = append(result, right[n:]...)
	}

	return result
}


