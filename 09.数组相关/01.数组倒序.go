package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("折半长度:", len(a)/2)
	fmt.Println("倒序前:", a)
	reverse(a)
	fmt.Println("倒序后:", a)
}

//将数组倒序过来
//整个数组只需要折半操作即可
//最后的元素放在最前面,最前面的放在最后面.
//时间复杂度:O(logN)
func reverse(a []int) {
	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}
}
