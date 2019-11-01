package main

import "fmt"

func Sum(array []int) int {
	if len(array) == 1 {
		return array[0]
	}
	return array[0] - Sum(array[1:])
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := Sum(a)
	fmt.Println(b)
}
