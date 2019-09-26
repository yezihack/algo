package recursive

import (
	"fmt"
	"testing"
)

//递归学习
//递归：无限调用自身这个函数，每次调用总会改动一个关键变量，直到这个关键变量达到边界的时候，不再调用。
//必须有规律才能使用递归操作
//每个递归函数都有两个部分组成:
//1. 基线条件 指函数不再调用自己,避免无限循环
//2. 递归条件 函数调用自己
//核心:　栈实现，压栈，出栈．
func countDown(i int) {
	fmt.Println(i)
	if i <= 1 { // 基线条件
		return
	}
	countDown(i - 1) //递归条件
}
func TestCountDown(t *testing.T) {
	countDown(10)
}

//累加, 使用递归算法
//每一次拿一个元素出来, 第二调用自己时,数组减少一个元素
func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + Sum(arr[1:])
}
func TestSum(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(a))
}

//计算数组数量,使用递归完成
// Count([1,2,3])
// 第1次: 1 + Count([2,3])
// 第2次: 1 + Count([3])
// 第3次: 1 + Count([])
func Count(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + Count(arr[1:])
}
func TestCount(t *testing.T) {
	arr := []int{1, 2, 4, 5, 8, 9}
	fmt.Println(Count(arr))
}

//非递归的二分查找
func BinarySearch(target int, arr []int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if target == arr[mid] {
			return mid
		} else if target > arr[mid] {
			low = mid + 1
		} else if target < arr[mid] {
			high = mid - 1
		}
	}
	return -1
}
func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 4, 5}
	fmt.Println(BinarySearch(1, arr))
	fmt.Println(BinarySearch(2, arr))
	fmt.Println(BinarySearch(4, arr))
	fmt.Println(BinarySearch(5, arr))
}

//递归二分查找
//二分的低位与高位, 每次做为参数传递到下一次调用的函数.
func BinarySearchRecursion(target int, arr []int, low, high int) int {
	mid := (low + high) / 2
	if low > high { //基线条件
		return -1
	}
	if target > arr[mid] {
		low = mid + 1
	} else if target < arr[mid] {
		high = mid - 1
	} else {
		return mid
	}
	return BinarySearchRecursion(target, arr, low, high) //递归条件
}

func TestBinarySearchRecursion(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(BinarySearchRecursion(2, arr, 0, len(arr)-1))
	fmt.Println(BinarySearchRecursion(5, arr, 0, len(arr)-1))
}

//找到最大的元素
//不断的缩小范围, 缩到数组只剩一个元素,则是最大的元素
//如果索引0比索引1大,则交换位置, 再调用自己缩小范围.
func Max(arr []int) int {
	if len(arr) == 1 { //基线条件
		return arr[0]
	}
	if arr[0] > arr[1] { //如果索引0比索引1大,则交换位置
		arr[0], arr[1] = arr[1], arr[0]
	}
	return Max(arr[1:]) //递归条件
}
func TestMax(t *testing.T) {
	arr := []int{1, 82, 2, 33, 71, 22}
	fmt.Println(Max(arr))
}

type ChildItem struct {
	Id       int
	Name     string
	ParentId int
	Child    []ChildItem
}

//二维数据转换成树结构
func ArrayToTree(list []ChildItem, pid int) []ChildItem {
	tree := make([]ChildItem, 0)
	for _, item := range list {
		fmt.Printf("id:%d, parent_id:%d, pid:%d\n", item.Id, item.ParentId, pid)
		if item.ParentId == pid {
			item.Child = ArrayToTree(list, item.Id)
			tree = append(tree, item)
		}
	}
	fmt.Println(tree)
	return tree
}
func TestArrayToTree(t *testing.T) {
	list := []ChildItem{
		{
			Id:       1,
			Name:     "北京",
			ParentId: 0,
		},
		{
			Id:       2,
			Name:     "朝阳",
			ParentId: 1,
		},
		{
			Id:       3,
			Name:     "天津",
			ParentId: 0,
		},
	}
	ret := ArrayToTree(list, 0)
	fmt.Println(ret)
}

//求最大公约数
func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
func TestGcd(t *testing.T) {
	fmt.Println(gcd(2, 10))
	fmt.Println(gcd(12, 3))
}

var step int

func move(n int, from, to byte) {
	step++
	fmt.Printf("第%d步:将%d号盘子从%c移动到%c\n", step, n, from, to)
}
func hanio(n int, start, mid, end byte) {
	if n == 1 {
		move(n, start, end)
	} else {
		hanio(n-1, start, end, mid)
		move(n, start, end)
		hanio(n-1, mid, start, end)
	}
}
func TestHanio(t *testing.T) {
	hanio(5, 'A', 'B', 'C')
}
