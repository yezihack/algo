package _6_排序算法

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"github.com/yezihack/algo/16.排序算法/exercise"
	"testing"
)

func TestInsertion(t *testing.T) {
	arr := []int{8, 2, 5, 4}
	fmt.Println(arr)
	Insertion(arr)
	fmt.Println(arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{3, 2, 5}
	fmt.Println(arr)
	InsertionSort(arr)
	fmt.Println(arr)
}

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
func TestBubble(t *testing.T) {
	arr := []int{8, 2, 3, 5, 7, 9}
	Bubble(arr)
	fmt.Println(arr)
}
//冒泡排序
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

func TestQuickSortSimple(t *testing.T) {
	arr := []int{5, 2, 8, 7, 4, 3, 1, 6}
	QuickSortSimple(arr, 0, len(arr)-1)
	fmt.Println(arr)
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
//选择排序
func TestSelectSort(t *testing.T) {
	arr := []int{3, 1, 4, 2}
	SelectSort(arr)
	src.Asset(1, t, 2, arr[1])
	src.Asset(2, t, 4, arr[3])
	src.Asset(3, t, 1, arr[0])
}
//归并排序
func TestMergeSort(t *testing.T) {
	fmt.Println(MergeSort([]int{8, 2, 7, 6, 3, 2, 1, 5}))
	fmt.Println(MergeSort([]int{3, 5, 2, 1, 4}))
}
//归并
func TestMergeSort2(t *testing.T) {
	fmt.Println(exercise.MergeSort([]int{8, 2, 7, 6, 3, 2, 1, 5}))
	fmt.Println(exercise.MergeSort([]int{3, 5, 2, 1, 4}))
}
//快排
func TestQuickSort2(t *testing.T) {
	arr := []int{8, 3, 2, 5, 4, 1, 7, 6}
	exercise.QuickSort(arr)
	fmt.Println(arr)
}
//桶排序
func TestBucketSort(t *testing.T) {
	arr := []int{8, 3, 2, 5, 4, 1, 7, 6, 3}
	BucketSort(arr)
	fmt.Println(arr)
}

