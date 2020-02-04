package LeetCode

import (
	"fmt"
	"sort"
	"testing"
)

//https://leetcode-cn.com/problems/relative-sort-array


//利用桶排序的升级版本,计数排序实现解题
//先将元素发送到桶里, 然后收集元素.
//对于此道题, 需要注意收集元素里, 我们需要先收集arr2数组的元素. 剩余部分再按升序收集

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	bucket := make([]int, 1001)
	newBucket := make([]int, len(arr1))
	//第一步:先将元素发送到各个桶里去
	//将arr1放在桶里面, 相同的元素则计数处理.
	for i := 0; i < len(arr1); i ++ {
		bucket[arr1[i]] ++
	}
	fmt.Println(bucket)
	//第二步:按arr2收集部分数据.然后将原桶里的数据减去掉.否则会出现重复.
	newCount := 0
	for i := 0; i < len(arr2); i ++ { //循环arr2
		for bucket[arr2[i]] > 0 { //查找bucket存在的元素出现的个数.直到为0,不再遍历
			newBucket[newCount] = arr2[i] //收集元素存放在新的桶里.
			newCount ++
			bucket[arr2[i]] -- //减去已经拿走的数据.
		}
	}
	//第三步: 将桶里剩余的元素全部按升序方式收集.如果是要求按倒序的话,就从最后的元素收集开始.
	for i := 0; i < len(bucket); i ++ {
		for bucket[i] > 0 {
			newBucket[newCount] = i //bucket是桶, 下标则是实际的值.
			newCount ++ //新桶下标递增
			bucket[i] -- //移除bucket桶已经拿走的数据.
		}
	}
	return newBucket
}

func TestRelativeSortArray(t *testing.T) {
	 r := RelativeSortArray([]int{2,3,1,3,2,4,6,7,9,2,19}, []int{2,1,4,3,9,6})
	 fmt.Println(r)
	r = relativeSortArray([]int{2,3,1,3,2,4,6,7,9,2,19}, []int{2,1,4,3,9,6})
	fmt.Println(r)
	arr1 := []int{8, 2, 1, 9, 10}
	fmt.Println(sort.IntSlice(arr1))
	sort.Sort(sort.IntSlice(arr1))
	fmt.Println(arr1)
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int, len(arr2))
	for _, num := range arr2 {
		m[num] = 0
	}
	l := len(arr1) - 1
	index := l
	for i := l; i >= 0; i-- {
		num := arr1[i]
		if _, ok := m[num]; !ok {
			arr1[index] = num
			index--
		} else {
			m[num] ++
		}
	}
	index = 0
	for _, num := range arr2 {
		for i := m[num] - 1; i >= 0; i-- {
			arr1[index] = num
			index ++
		}
	}
	fmt.Println(arr1)
	sort.Sort(sort.IntSlice(arr1[index:]))

	return arr1
}