package sort

import (
	"fmt"
	"testing"
)

//填空法(i,j指针必须j指针先开始)
//首先找最左边的值做基准值:p, 然后i指针代表最左边, j指针代表最右边
//i指针每次向右移与p值比较,如果i < p则右移,否则停止
//j指针每次向左移与p值比较,如果j > p则左移,否则停止
//i,j指针都停止后再 i与j的值交换
//再进行一次循环.直到i与j重合, 然p值填到i与j重合的位置.
//此时的数组,以p为准,p的左边都是小于p的值,p的右边都大于p的值.
//然后分别将p左边的数组排序与p右边的数组排序.
//以i为基准, i 与 left 未重合则代表,数组left:i之间数组个数大于1.(数组只有一个元素则无需排序啦)
func QSortFill(list []int, left, right int) {
	pivot := list[left]
	i, j := left, right
	for i < j {
		for i < j && list[j] > pivot {
			j--
		}
		list[i] = list[j]
		for i < j && list[i] < pivot {
			i++
		}
		list[j] = list[i]
	}
	list[i] = pivot
	if i-left > 1 {
		QSortFill(list, left, i-1)
	}
	if right-i > 1 {
		QSortFill(list, i+1, right)
	}
}
func TestQSortFill(t *testing.T) {
	list := []int{8, 2, 3, 5, 9, 1}
	fmt.Println("unsort:", list)
	QSortFill(list, 0, len(list)-1)
	fmt.Println("sorted:", list)
}

//指针交换法(i,j指针无轮哪个先走都可以)
//先找到一个基准值p, 然后i指针代表最左边, j指针代表最右边
//i指针每次向右移与p值比较,如果i < p则右移,否则停止
//j指针每次向左移与p值比较,如果j > p则左移,否则停止
//i,j指针都停止后再 i与j的值交换
//再进行一次循环.直到i与j重合, 然p值填到i与j重合的位置.
//此时的数组,以p为准,p的左边都是小于p的值,p的右边都大于p的值.
//然后分别将p左边的数组排序与p右边的数组排序.
//以i为基准, i 与 left 未重合则代表,数组left:i之间数组个数大于1.(数组只有一个元素则无需排序啦)
func QSortPoint(list []int, left, right int) {
	pivot := list[left]
	i, j := left, right
	for i < j {
		for i < j && list[j] > pivot {
			j--
		}
		for i < j && list[i] < pivot {
			i++
		}
		if i < j {
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i] = pivot
	if i-left > 1 {
		QSortPoint(list, left, i-1)
	}
	if right-i > 1 {
		QSortPoint(list, i+1, right)
	}
}
func TestQSortPoint(t *testing.T) {
	list := []int{8, 3, 2, 1, 5, 11, 21, 55, 9}
	fmt.Println("QSortPoint排序前:", list)
	QSortPoint(list, 0, len(list)-1)
	fmt.Println("QSortPoint排序后:", list)
}

//QSortPoint的改良版本. 只要left大于等于right就不再递归.方便记忆
func QSortPoint2(list []int, left, right int) {
	if left >= right {
		return
	}
	pivot := list[left]
	i, j := left, right
	for i < j {
		for i < j && list[j] > pivot {
			j--
		}
		for i < j && list[i] < pivot {
			i++
		}
		if i < j {
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i] = pivot
	QSortPoint(list, left, i-1)
	QSortPoint(list, i+1, right)
}
func TestQSortPoint2(t *testing.T) {
	list := []int{8, 3, 2, 1, 5, 9}
	fmt.Println("TestQSortPoint2排序前:", list)
	QSortPoint2(list, 0, len(list)-1)
	fmt.Println("TestQSortPoint2排序后:", list)
}

func QSortBook(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		for values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}

	}
	values[p] = temp
	if p-left > 1 {
		QSortBook(values, left, p-1)
	}
	if right-p > 1 {
		QSortBook(values, p+1, right)
	}
}
func TestQSortBook(t *testing.T) {
	list := []int{8, 3, 2, 1, 5, 9}
	fmt.Println(list)
	QSortBook(list, 0, len(list)-1)
	fmt.Println(list)
}

func TestFastSort(t *testing.T) {
	list := []int{8, 3, 2, 1, 5, 9}
	fmt.Println(list)
	FastSort(list)
	fmt.Println(list)
}