package LeetCode

import "fmt"

//map[2:7 5:4 16:7 17:1 20:6 26:3]  2和16都出现7次,就不是独一无二
//时间复杂度O(NlogN),空间复杂度:O(n)
//思路: 先使用map进行数组出现的元素进行统计个数, 然后将出现的次数存放到一个数组里,
// 然后再进行快排, 再进行快慢指针判断是否存在重复的元素.
func UniqueOccurrences(arr []int) bool {
	if len(arr) == 0 || len(arr) == 1 {
		return false
	}
	have := make(map[int]int)
	for _, v := range arr {
		if _, ok := have[v]; ok {
			have[v] ++
		} else {
			have[v] = 1
		}
	}
	lst := make([]int, len(have))
	i := 0
	for _, cnt := range have {
		lst[i] = cnt
		i ++
	}

	quickSort(lst, 0, len(lst)-1)
	//fmt.Println(lst)
	for i := 0; i < len(lst) - 1;i ++ {
		//fmt.Printf("arr[%d]=%d, arr[%d]=%d\n", i, lst[i], i+1, lst[i+1])
		if lst[i] == lst[i+1] {
			return false
		}
	}
	return true
}
func quickSort(arr []int, left, right int) {
	pivot := arr[left]
	i, j := left, right
	for i < j {
		for i < j && pivot <= arr[j] {
			j --
		}
		arr[i] = arr[j]
		for i < j && pivot >= arr[i] {
			i ++
		}
		arr[j] = arr[i]
	}
	arr[i] = pivot
	if i - left > 1 {
		quickSort(arr, left, i - 1)
	}
	if right - i > 1 {
		quickSort(arr, i + 1, right)
	}
}
//别人的思路
//借用二个map实现
//第一个map统计数字的次数.
//第二个map再判断次数里是否出现重复的次数.若出现则说明,有同一样
func UniqueOccurrences2(arr []int) bool {
	mapData := make(map[int]int)
	for _, v := range arr {
		mapData[v]++
	}
	fmt.Println(mapData)
	arrData := make(map[int]int)
	for _, v := range mapData {
		arrData[v]++
	}
	fmt.Println(arrData)
	for _, v := range arrData {
		if v != 1 {
			return false
		}
	}
	return true

}