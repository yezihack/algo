package _4_蓝蝙蝠面试算法

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

/*
数组转换成最小值
如arr := []int{16, 15, 27, 9}
输出: 1516279
*/

//冒泡方法, 时间复杂度O(n^2)
func PrintMinNumberByBubbling(numbers []int) string {
	length := len(numbers)

	for i := 0; i < length; i ++ {
		for j := 0; j < length-1-i; j++ {
			if cmp(numbers[j], numbers[j+1]) {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	build := strings.Builder{}
	for _, v := range numbers {
		build.WriteString(strconv.Itoa(v))
	}
	return build.String()
}

//利用系统自带排序函数
func MinNumberBySort(numbers []int) string {
	if len(numbers) == 0 {
		return ""
	}
	sort.Slice(numbers, func(i, j int) bool {
		return !cmp(numbers[i], numbers[j])
	})
	build := strings.Builder{}
	for _, v := range numbers {
		build.WriteString(strconv.Itoa(v))
	}
	return build.String()
}
//快排实现
func MinNumbersByQuickSort(numbers []int) string {
	qSort(numbers, 0, len(numbers)-1)
	build := strings.Builder{}
	for _, v := range numbers {
		build.WriteString(strconv.Itoa(v))
	}
	return build.String()
}

//快排
func qSort(arr []int, left, right int) {
	i, j := left, right
	pivot := arr[left]
	for left < right {
		for left < right && cmp(arr[right], pivot) {//比较大小, 大于中间值的不移动,反之移动.
			right --
		}
		arr[left] = arr[right]
		for left < right && cmp(pivot, arr[left]) {//比较大小, 小于中间值的不移动.反之移动
			left ++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	if left - i > 1 {
		qSort(arr, i, left-1)
	}
	if j - left > 1 {
		qSort(arr, left+1, j)
	}
}

//核心函数
// 如 9, 12. 我们需要做912与129的比较, 显然912>129
//如果ab < ba 的话则返回true,否则返回false
func cmp(a, b int) bool {
	if a == b {
		return false
	}
	sa := strconv.Itoa(a)
	sb := strconv.Itoa(b)
	sa2 := sa + sb
	sb2 := sb + sa
	ma, err := strconv.Atoi(sa2)
	if err != nil {
		log.Fatal("a params transfer failed")
		return false
	}
	nb, err := strconv.Atoi(sb2)
	if err != nil {
		log.Fatal("b params transfer failed")
		return false
	}
	fmt.Printf("m:%d,n:%d\n", ma,nb)
	if ma > nb { //相当于mn > nm 将最小的元素置换到最前面.如果是小于则最大元素置换到最前面.
		return true
	}
	return false
}