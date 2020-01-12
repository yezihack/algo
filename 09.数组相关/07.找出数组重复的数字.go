package main

import "fmt"

/*
指剑Offer第二版的第3题.
在一个长度为n的数组里所有的数字都在0~n-1的范围内.
数组中某些数字是重复的.
例: 长度为7的数组[2, 3, 1, 0, 2, 5, 3] 重复数字: 2, 3

我觉得要是
*/

//第一种方式,暴力方法
//每一个元素进行查看一遍,O(n^2)
func repeat(arr []int) []int {
	rep := make([]int, 0)
	for i := 0; i < len(arr); i ++ {
		for j := i + 1; j < len(arr); j ++ {
			if arr[i] == arr[j] {
				rep = append(rep, arr[i])
				break
			}
		}
	}
	return rep
}
//第二种, 借用一个数组,hash方法
//建一个数组,最长度就是数组里最大元素的个数.然后循环一遍Arr,
// 把value当做key存入hashArray数组里,如果存在则说明这个元素是重复的
//时间复杂度O(n), 空间O(n)
func duplicateByHash(arr []int) []int{
	max := 0
	for i := 0; i < len(arr); i ++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	hashArray := make([]int, max+1)
	for i := 0; i < max+1; i ++ {
		hashArray[i] = -1
	}
	rep := make([]int, 0)
	for i := 0; i < len(arr); i ++ {
		if hashArray[arr[i]] >= 0 {
			rep = append(rep, arr[i])
		} else {
			hashArray[arr[i]] = 0
		}
	}
	return rep
}
//找到下标与值一一对应, 如果不相同则说明重复.
// 有个缺点, 如果值的大小与下标本来不一致,这种算法就不行.
func duplicate3(arr []int) []int {
	rep := make([]int, 0)
	for i := 0; i < len(arr); i ++ {
		for arr[i] != i {
			tmp := arr[i] //数组的值,需要找到与值相同的下标位置
			if arr[i] == arr[tmp] {
				rep = append(rep, arr[i])
				break
			}
			arr[i] = arr[tmp] //下标与值对应起来
			arr[tmp] = tmp //下标与值对应起来
		}
	}
	return rep
}
//先使用快排, O(logN), 然后使用数组第一个元素与第二元素相互比较一下
//如果相等则说明是重复的值.
func replicateByQuick(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	Quick(arr, 0, len(arr)-1)

	rep := make([]int, 0)
	tmp := 0
	for i := 1; i < len(arr)-1; i ++ {
		if arr[i] == arr[tmp] {
			rep = append(rep, arr[i])
		} else {
			tmp = i
		}
	}
	return rep
}
func Quick(arr []int, left, right int) {
	if left < right {
		index := partition(arr, left, right)
		Quick(arr, left,  index-1)
		Quick(arr, index + 1, right)
	}
}
func partition(arr []int, left, right int) int {
	pivot := arr[left]
	i, j := left, right
	for i < j {
		for i < j && arr[j] >= pivot {
			j --
		}
		arr[i] = arr[j]
		for i < j && arr[i] <= pivot {
			i ++
		}
		arr[j] = arr[i]
	}
	arr[i] = pivot
	return i
}

func main() {
	fmt.Println(repeat([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(duplicateByHash([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(duplicate3([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(replicateByQuick([]int{2, 3, 1, 2, 2, 5, 3}))
}