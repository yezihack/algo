package _6_排序算法

import (
	"fmt"
	"testing"
)


//插入排序,不带哨兵
func TestInsertSort(t *testing.T) {
	arr := []int{8, 2, 1, 7, 4, 3, 5, 6}
	fmt.Println("before", arr)
	InsertSort(arr)
	fmt.Println("after", arr)
	if arr[0] != 1 {
		t.Errorf("expect:%d, actual:%d\n", 1, arr[0])
	}
	if arr[len(arr)-1] != 8 {
		t.Errorf("expect:%d, actual:%d\n", 8, arr[len(arr)-1])
	}
	if arr[3] != 4 {
		t.Errorf("expect:%d, actual:%d\n", 4, arr[3])
	}
}
//插入排序,带哨兵
func TestInsertSortSentry(t *testing.T) {
	arr := []int{0, 2, 1, 7, 4, 3, 6}
	fmt.Println("before", arr)
	InsertSortSentry(arr)
	fmt.Println("after", arr, len(arr))
	if arr[1] != 1 {
		t.Errorf("expect:%d, actual:%d\n", 1, arr[1])
	}
	if arr[len(arr)-1] != 7 {
		t.Errorf("expect:%d, actual:%d\n", 7, arr[len(arr)-1])
	}
	if arr[3] != 3 {
		t.Errorf("expect:%d, actual:%d\n", 3, arr[3])
	}
}
//折半插入排序
func TestBInsertSort(t *testing.T) {
	arr := []int{8, 2, 1, 7, 4, 3, 5, 6}
	fmt.Println("before", arr)
	BInsertSort(arr)
	fmt.Println("after", arr)
	if arr[0] != 1 {
		t.Errorf("expect:%d, actual:%d\n", 1, arr[0])
	}
	if arr[len(arr)-1] != 8 {
		t.Errorf("expect:%d, actual:%d\n", 8, arr[len(arr)-1])
	}
	if arr[3] != 4 {
		t.Errorf("expect:%d, actual:%d\n", 4, arr[3])
	}
}
//折半插入排序,带哨兵
func TestBSentryInsertSort(t *testing.T) {
	arr := []int{0, 2, 1, 7, 4, 3, 5, 6}
	fmt.Println("before", arr)
	BSentryInsertSort(arr)
	fmt.Println("after", arr)
	if arr[1] != 1 {
		t.Errorf("expect:%d, actual:%d\n", 1, arr[1])
	}
	if arr[len(arr)-1] != 7 {
		t.Errorf("expect:%d, actual:%d\n", 7, arr[len(arr)-1])
	}
	if arr[3] != 3 {
		t.Errorf("expect:%d, actual:%d\n", 3, arr[3])
	}
}
//希尔排序
func TestShellSort(t *testing.T) {
	arr := []int{8, 2, 1, 7, 4, 3, 5, 6}
	fmt.Println("before", arr)
	ShellSort(arr)
	fmt.Println("after", arr)
	if arr[0] != 1 {
		t.Errorf("expect:%d, actual:%d\n", 1, arr[0])
	}
	if arr[len(arr)-1] != 8 {
		t.Errorf("expect:%d, actual:%d\n", 8, arr[len(arr)-1])
	}
	if arr[3] != 4 {
		t.Errorf("expect:%d, actual:%d\n", 4, arr[3])
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{8, 2, 1, 7, 4, 3, 5, 6}
	arr = []int{1,2,3,4,5,6,7,8}
	fmt.Println("before", arr)
	BubbleSort(arr)
	fmt.Println("after", arr)
	if arr[0] != 1 {
		t.Errorf("expect:%d, actual:%d\n", 1, arr[0])
	}
	if arr[len(arr)-1] != 8 {
		t.Errorf("expect:%d, actual:%d\n", 8, arr[len(arr)-1])
	}
	if arr[3] != 4 {
		t.Errorf("expect:%d, actual:%d\n", 4, arr[3])
	}
}
//冒泡改进
func TestBubbleSortFlag(t *testing.T) {
	arr := []int{8, 2, 1, 7, 4, 3, 5, 6}
	fmt.Println("before", arr)
	BubbleSortFlag(arr)
	fmt.Println("after", arr)
	if arr[0] != 1 {
		t.Errorf("expect:%d, actual:%d\n", 1, arr[0])
	}
	if arr[len(arr)-1] != 8 {
		t.Errorf("expect:%d, actual:%d\n", 8, arr[len(arr)-1])
	}
	if arr[3] != 4 {
		t.Errorf("expect:%d, actual:%d\n", 4, arr[3])
	}
}
//快速排序,版本一(二个函数构成)
func TestFastSort(t *testing.T) {
	arr := []int{5, 2, 8, 7, 4, 3, 1, 6}
	fmt.Println("before", arr)
	FastSort(arr, 0, len(arr) - 1)
	fmt.Println("after", arr)
	if arr[0] != 1 {
		t.Errorf("expect:%d, actual:%d\n", 1, arr[0])
	}
	if arr[len(arr)-1] != 8 {
		t.Errorf("expect:%d, actual:%d\n", 8, arr[len(arr)-1])
	}
	if arr[3] != 4 {
		t.Errorf("expect:%d, actual:%d\n", 4, arr[3])
	}
}
//快速排序,版本二(一个函数)
func TestQuickSort(t *testing.T) {
	arr := []int{5, 2, 8, 7, 4, 3, 1, 6}
	fmt.Println("before", arr)
	QuickSort(arr, 0, len(arr) - 1)
	fmt.Println("after", arr)
	if arr[0] != 1 {
		t.Errorf("expect:%d, actual:%d\n", 1, arr[0])
	}
	if arr[len(arr)-1] != 8 {
		t.Errorf("expect:%d, actual:%d\n", 8, arr[len(arr)-1])
	}
	if arr[3] != 4 {
		t.Errorf("expect:%d, actual:%d\n", 4, arr[3])
	}
}

//快速排序版本3
func TestQuickSort3(t *testing.T) {
	arr := []int{5, 2, 8, 7, 4, 3, 1, 6}
	fmt.Println("before", arr)
	QuickSort3(arr)
	fmt.Println("after", arr)
	if arr[0] != 1 {
		t.Errorf("expect:%d, actual:%d\n", 1, arr[0])
	}
	if arr[len(arr)-1] != 8 {
		t.Errorf("expect:%d, actual:%d\n", 8, arr[len(arr)-1])
	}
	if arr[3] != 4 {
		t.Errorf("expect:%d, actual:%d\n", 4, arr[3])
	}
}
//简单选择排序 不稳定排序,O(N^2)
func TestSimplySelectSort(t *testing.T) {
	fmt.Println(1<<1, 1<<2, 1<<3, 1<<4, 1<<5, 1<<6)
	fmt.Println(100>>1, 100>>2, 100>>3, 100>>4, 100>>5, 100>>6)
	arr := []int{5, 2, 8, 7, 4, 3, 1, 6}
	fmt.Println("before", arr)
	SimplySelectSort(arr)
	fmt.Println("after", arr)
	if arr[0] != 1 {
		t.Errorf("expect:%d, actual:%d\n", 1, arr[0])
	}
	if arr[len(arr)-1] != 8 {
		t.Errorf("expect:%d, actual:%d\n", 8, arr[len(arr)-1])
	}
	if arr[3] != 4 {
		t.Errorf("expect:%d, actual:%d\n", 4, arr[3])
	}
}







