package _9_查找算法

import (
	"fmt"
	"math"
)

//二分查找: 每一次查找都是折半查找.
//需要注意三点: left <= right
// low + (hight-low)/2 或者 (low+high)>>1
// low=mid+1, high=mid-1
func BinarySearch(arr []int, value int) int {
	if len(arr) == 0 {
		return -1
	}
	left, right := 0, len(arr) - 1
	mid := 0
	for left <= right {
		mid = (left + right) >> 1
		if value == arr[mid] {
			return mid
		} else if arr[mid] > value {
			right = mid - 1
		} else if arr[mid] < value {
			left = mid + 1
		}
	}
	return -1
}
//递归实现二分查找
func BReserveSearch(arr []int, low, high, value int) int {
	if low > high {
		return -1
	}
	mid := (low + high) >> 1
	if arr[mid] == value {
		return mid
	} else if arr[mid] > value {
		return BReserveSearch(arr, low, mid-1, value)
	} else {
		return BReserveSearch(arr, mid+1, high, value)
	}
}
//利用二分查找求平方根
func BSqrt(x float64) float64 {
	var (
		low float64 = 0
		high float64 = x
		mid float64
	)
	if x > 0 {
		low = 1
		high = x
	} else {
		low = x
		high = 1
	}
	mid = low + (high - low) / 2
	for float64(math.Abs(mid * mid - x)) > 0.000001 {
		if mid * mid < x {
			low = mid
		} else {
			high = mid
		}
		mid = low + (high - low) / 2
	}
	return mid
}
//数组里出现多个重复的数字,获取最前面的数字下标,如1,2,2,3我们要找2,1后面的2而不是3前面的2
func BFirstSearch(arr []int, value int) int {
	low, high := 0, len(arr)-1
	mid := 0
	for low <= high {
		mid = low + (high - low) >> 1
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			//如果是最左边的数,再向前找一个位置一定不是当前要找的值,否则high=mid-1继续向左找
			if mid == 0 || arr[mid-1] != value {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
//数组里有多少重复的数,查找到最后一个数字.如1,2,2,3.我们要找到3前面的2,不是1后面的2
func BLastSearch(arr []int, value int) int {
	low, high := 0, len(arr) - 1
	mid := 0
	for low <= high {
		mid = low + (high - low) >> 1
		if arr[mid] < value {
			low = mid + 1
		} else if arr[mid] > value {
			high = mid - 1
		} else {
			fmt.Println("mid", mid, "low", low, "high", high)
			if mid == len(arr) - 1 || arr[mid + 1] != value {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

//查找第一个大于等于给定值的元素
func BGtSearch(arr []int, value int) int {
	low, high := 0, len(arr)-1
	mid := 0
	for low <= high {
		mid = low + (high-low)>>1
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			if mid == len(arr) - 1 || arr[mid + 1] > value {
				return mid + 1
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}
//查找最后一个小于等于给定值的元素
func BLTSearch(arr []int, value int) int {
	low, high := 0, len(arr) - 1
	mid := 0
	for low <= high {
		mid = low + (high - low) >> 1
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] < value {
				return mid - 1
			} else {
				high = mid -1
			}
		}
	}
  	return -1
}
