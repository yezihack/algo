package sort

import (
	"fmt"
	"testing"
)

func QuickSort(list []int, low, high int) {
	//记录开始与结束的位置用于判断是否需要继续下一轮的排序.
	start, end := low, high
	//每次取中间的值做为基准值.
	mid := list[(low+high)/2]
	//低位点与高位点必须重合就结束一轮的排序.
	for low < high {
		//如果高位点大于基准值,则一直向左偏移
		for list[high] > mid {
			high--
		}
		//如果低位点小于基准值,则一直向右偏移
		for list[low] < mid {
			low++
		}
		//低位与高位停止偏移,表示需要交换数据位置,并高低位都偏移一次.否则会出现死循环.
		if low <= high { //必须保证是low <= high可以重合, 如果low,high偏移交叉了不进行偏移.否则会出错.
			list[low], list[high] = list[high], list[low]
			low++
			high--
		}
	}
	//如果开始位置还小于向左偏移的高位点的话,表示高位还没有完成偏移量,也就是说每一次的切隔,直到最高位偏移到只剩一个元素.
	if start < high {
		QuickSort(list, start, high)
	}
	//同上
	if end > low {
		QuickSort(list, low, end)
	}
}
func TestQuickSort(t *testing.T) {
	list := []int{2, 3, 1, 4, 5, 8, 6, 7}
	fmt.Println("排序前:", list)
	QuickSort(list, 0, len(list)-1)
	fmt.Println("排序后:", list)
}

func Quick(list []int) {
	if len(list) <= 1 { //一个元素, 无需排序
		return
	}
	key := list[len(list)-1] //每次找第一个做基准值
	low := 0                 //开始的位置
	high := len(list) - 1    //结束的位置
	for low < high {         //高位与低位不重合,则行判断交换
		if low < high && list[low] < key { //如果低位小于基准值, 则向右偏移
			low++
		}
		list[high] = list[low]              //停止偏移则将低位的数据交换到高位.
		if low < high && list[high] > key { //如果高位大于基准值,则向左偏移
			high--
		}
		list[low] = list[high] //停止偏移则将高位的数据交换到低位
	}
	list[low] = key //将基准值放值低位下标处.
	//fmt.Printf("privot:%d, low:%d, 左边:%v, 右边:%v\n", key, low, list[:low], list[low+1:])
	Quick(list[:low])
	Quick(list[low+1:])
}
func TestQuick(t *testing.T) {
	list := []int{2, 3, 1, 4, 5, 8, 6, 7}
	fmt.Println("排序前:", list)
	Quick(list)
	fmt.Println("排序后:", list)
}
