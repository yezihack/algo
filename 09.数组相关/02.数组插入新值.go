package main

import (
	"code.qschou.com/golang/errors"
	"fmt"
)

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	fmt.Println("ori", a)
	Insert(a, 2, 22)
	Delete(a, 2)
}
func Insert(a [5]int, index, value int) error {
	//1.判断位置是否合法
	if index < 0 || index > len(a)-1 {
		return errors.New("Index is illegal")
	}
	//2.判断是否满了.
	//移动index以后的位置,插入新value
	for j := len(a) - 1; j >= index; j-- {
		a[j] = a[j-1]
	}
	a[index] = value
	fmt.Println(a)
	return nil
}
func Delete(a [5]int, index int) error {
	if index < 0 || index > len(a) {
		return errors.New("Index is illegal")
	}
	for i := index; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a[len(a)-1] = 0
	fmt.Println(a)
	return nil
}
