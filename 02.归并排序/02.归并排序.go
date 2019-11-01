package main

import "fmt"

func Merge(a []int, left, mid, right int) {
	temp := make([]int, len(a))
	i := left    //左边部分指针
	j := mid + 1 //右边部分数组指针
	t := 0       //临时存放归并之后的数组
	fmt.Printf("i:%d, j:%d, t:%d, mid:%d, left:%d,right:%d\n", i, j, t, mid, left, right)
	//第一步:
	//当i的开始位置一直移动到中间位置mid就算完成了左边的移动
	//反之,当j的开始位置一直移动到最右边位置right就算完成了右边的移动.
	//这里会存在左边移完了, 右边还剩余. 反之右边移完了, 左边还剩余.
	for i <= mid && j <= right {
		if a[i] < a[j] {
			temp[t] = a[i]
			t++
			i++
		} else {
			temp[t] = a[j]
			t++
			j++
		}
	}
	//第二步：解决未移动完的一边情况
	//解决: 右边移完了, 左边还剩余
	for i <= mid {
		temp[t] = a[i]
		i++
		t++
	}
	//解决: 左边移完了, 右边还剩余
	for j <= right {
		temp[t] = a[j]
		j++
		t++
	}
	//第三步: 将临时temp数组的值移动到a数组里去
	t = 0
	for t <= right {
		a[t] = temp[t]
		t++
	}
}
func main() {
	a := []int{1, 3, 5, 6, 2, 7, 8, 9}
	fmt.Println("排序前:", a)
	mid := len(a)/2 - 1
	Merge(a, 0, mid, len(a)-1)
	fmt.Println("排序后:", a)
}
