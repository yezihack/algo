package exercise

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2 //将数组一分为二
	left := MergeSort(arr[0:mid])//拆分前半部分
	right := MergeSort(arr[mid:])//拆分后半部分
	return merge(left, right) //将拆分的再合并
}
//将有序数组合并一个有序数组
func merge(left, right []int) []int {
	temp := make([]int, 0)//申请临时空间存储合并后的数组
	m, n := 0, 0 //申请两个指针
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] < right[n] {
			temp = append(temp, left[m])
			m ++
		} else {
			temp = append(temp, right[n])
			n ++
		}
	}
	if m < l {
		temp = append(temp, left[m:]...)
	}
	if n < r {
		temp = append(temp, right[n:]...)
	}
	return temp
}