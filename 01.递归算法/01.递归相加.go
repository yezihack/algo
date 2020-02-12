package Recurive

import "fmt"

func Sum(array []int) int {
	if len(array) == 1 {
		return array[0]
	}
	return array[0] + Sum(array[1:])
}

func SumV2(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	fmt.Println(arr)
	return arr[0] + SumV2(arr[1:])
}

