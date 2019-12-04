package exercise

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}
//将有序数组合并一个有序数组
func merge(left, right []int) []int {
	temp := make([]int, 0)
	m, n := 0, 0
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