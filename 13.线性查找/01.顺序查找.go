package main

import "fmt"

type KeyType int
func main(){
	a := []int{1,2,3, 5, 4, 7, 6}
	fmt.Println(SearchSeq(3, a))
}
//增加一个哨兵,加入到数组这一个元素里.
//防止数组越界.
func SearchSeq(key int, a []int) int {
	if a[0] == key {
		return 0
	}
	a[0] = key
	i := len(a) -1
	for a[i] != key {
		i --
	}
	return i
}