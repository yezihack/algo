package dichotomy

import (
	"fmt"
	"math"
	"sort"
)

func dichotomy(val int, lst []int) int {
	//进行排序 1, 2, 3, 4, 5, 6,
	sort.Ints(lst)
	var start, end = 0, len(lst) - 1 //设置最大下标值,注意计算mid时需要加1, 原因下面有说
	var mid = (start + end) >> 1
	if val == lst[start] {
		return start
	}
	if val == lst[end-1] {
		return end - 1
	}
	if val > lst[end] || val < lst[start] {
		return -1
	}
	for start < end {
		fmt.Printf("start:%d, end:%d, mid:%d, want: %d \n", start, end, mid, val)
		if val == lst[mid] { //如果等于中间值,则返回
			return mid
		} else if val < lst[mid] { //如果目标值小于中间值,则将中间下标赋值给end,缩小范围查找
			end = mid
		} else { //如果目标值大于中间值,则将中间下标赋值给start,缩小范围查找
			start = mid
		}
		//这里为什么加1呢? 当下标最后2位相加除以2,如1,2,3,4,5,
		// 下标是3+4=7, 7/2 = 3小于最后下标4, 也就是说会出现死循环.所以加1
		mid = (start+end)/2 + 1
	}
	return -1
}

//第二种写法
func dichotomy2(val int, lst []int) int {
	var start, end = 0, len(lst) //求出最小值和最大数
	var mid = (start + end) >> 1
	if val == lst[start] {
		return start
	}
	if val == lst[end-1] {
		return end - 1
	}
	if val > lst[end-1] || val < lst[start] {
		return -1
	}
	for start < end-1 { //end -1 很重要.否则会死循环
		fmt.Printf("start:%d, end:%d, mid:%d, want: %d \n", start, end, mid, val)
		if val == lst[mid] { //如果等于中间值,则返回
			return mid
		} else if val < lst[mid] { //如果目标值小于中间值,则将中间下标赋值给end,缩小范围查找
			end = mid
		} else { //如果目标值大于中间值,则将中间下标赋值给start,缩小范围查找
			start = mid
		}
		mid = (start + end) / 2 //求出中位下标
	}
	return -1
}

//推荐写法
func dichotomy3(val int, lst []int) int {
	//实例三个值, 用于缩小范围,快速查找
	//start=0 代表下标为0
	//end=len(lst) -1 代表数组最大下标
	//mid=初使为0
	start, end, mid := 0, len(lst)-1, 0
	if val == lst[start] {
		return start
	}
	if val == lst[end-1] {
		return end - 1
	}
	if val > lst[end] || val < lst[start] {
		return -1
	}
	for {
		mid = (start + end) >> 1                                                   //求中间下位值, golang整形除取最小值,相当于math.Floor
		fmt.Printf("start:%d, end:%d, mid:%d, want: %d \n", start, end, mid, val) //打印
		if val > lst[mid] {                                                       //如果目标值大于中间值,则将中间mid + 1赋值给start
			start = mid + 1
		} else if val < lst[mid] { //如果目标值小于中间值,则将中间mid-1赋值给end
			end = mid - 1
		} else { //表示值相等,找到目标下标
			break
		}
		if start > end { //如果start大于end则跳出,未找到
			mid = -1
			break
		}
	}
	return mid
}
func Find(k int, lst []int) int {
	sort.Ints(lst)
	left, right, mid := 1, len(lst), 0
	if k == lst[left-1] { //是否等于第一个值
		return left - 1
	}
	if k == lst[right-1] { //是否等于最后一个值
		return right - 1
	}
	if k > lst[right-1] || k < lst[left-1] {
		return -1
	}

	for {
		// mid向下取整
		mid = int(math.Floor(float64((left + right) / 2)))
		if lst[mid] > k {
			// 如果当前元素大于k，那么把right指针移到mid - 1的位置
			right = mid - 1
		} else if lst[mid] < k {
			// 如果当前元素小于k，那么把left指针移到mid + 1的位置
			left = mid + 1
		} else {
			// 否则就是相等了，退出循环
			break
		}
		// 判断如果left大于right，那么这个元素是不存在的。返回-1并且退出循环
		if left > right {
			mid = -1
			break
		}
	}
	// 输入元素的下标
	return mid
}
