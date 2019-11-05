package main

import "fmt"

//需要借用一个新的数组存放合并的值
//给A,B数组各新建一个实例指针0,逐个比较,将最小的值移动到新数组C里去.
//最后再次判断哪个数组没有移动完,逐个移完.
//产生一个合并后的数组.
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [6]int{5, 6, 7, 8, 9, 10}
	c := Merge(a, b)
	fmt.Println("merge array:", c)
}
func Merge(a [5]int, b [6]int) [11]int {
	var c [11]int
	aIndex, bIndex, cIndex := 0, 0, 0
	aLast, bLast := len(a)-1, len(b)-1
	for aIndex <= aLast && bIndex <= bLast {
		if a[aIndex] <= b[bIndex] {
			c[cIndex] = a[aIndex]
			aIndex++
		} else {
			c[cIndex] = b[bIndex]
			bIndex++
		}
		cIndex++
	}
	//查看哪个数组还没有移动完
	for aIndex <= aLast {
		c[cIndex] = a[aIndex]
		aIndex++
		cIndex++
	}
	for bIndex <= bLast {
		c[cIndex] = b[bIndex]
		bIndex++
		cIndex++
	}
	return c
}
