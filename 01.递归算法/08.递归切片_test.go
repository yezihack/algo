package Recurive

import (
	"fmt"
	"testing"
)

func TestGetOne(t *testing.T) {
	a := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println("result", GetOne(a))
}
