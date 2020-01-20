package _6_排序算法

import "fmt"

//冒泡排序
//两两比对,进行交换

func Bubble(arr []int) {
	var flag = false
	for i := 0; i < len(arr); i ++ {
		flag = false
		for j := 0; j < len(arr) - i - 1; j ++ {
			if arr[j] > arr[j + 1] { //比较
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true //表示有数据交换
			}
		}
		if !flag {
			break
		}
	}
}

//标准版本
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i ++ {
		for j := 0; j < len(arr)-i-1; j ++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
//冒泡改进
func BubbleSortFlag(arr []int) {
	x := 0
	for i := 0; i < len(arr); i ++ {
		flag := false
		for j := 0; j < len(arr) - i - 1; j ++ {
			if arr[j] > arr[j+1] {
				flag = true
				x ++
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !flag {
			break
		}
	}
	fmt.Print("BubbleSort", x)
}

func MaoPao(arr []int) {
	for i := 0; i < len(arr); i ++ {
		for j := 0; j < len(arr)-i-1;j ++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}