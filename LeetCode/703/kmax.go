package _03

import (
	"sort"
)
//设计一个找到数据流中第K大元素的类（class）。注意是排序后的第K大元素，不是第K个不同的元素。
type KthLargest struct {
	nums []int
	k int
}

func Constructor(k int, nums []int) KthLargest {
	l := len(nums)
	if k > l {
		panic("k > l")
	}
	arr := make([]int, l)
	sort.Ints(nums)
	arr = nums[l-k:]
	return KthLargest{
		nums:arr,
		k:k,
	}
}

func (this *KthLargest) Add(val int) int {
	if this.nums[len(this.nums) - 1] > val {
		return this.nums[0]
	}
	this.nums[0] = val
	sort.Ints(this.nums)
	return this.nums[0]
}
//实现排序.
func Sort(arr []int) {
	_Sort(arr, 0, len(arr) - 1)
}
func _Sort(arr []int, left, right int) {
	i, j := left, right
	pivot := arr[left]
	for i < j {
		for i < j && arr[j] > pivot {
			j --
		}
		arr[i] = arr[j]
		for i < j && arr[i] < pivot {
			i ++
		}
		arr[j] = arr[i]
	}
	arr[i] = pivot
	if i - left > 1 {
		_Sort(arr, left, i - 1)
	}
	if right - i > 1 {
		_Sort(arr, i + 1, right)
	}
}