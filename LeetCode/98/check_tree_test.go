package _8

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestCheckSortAsc(t *testing.T) {
	src.Asset(1, t, true, CheckSortAsc([]int{1,2,3,4}))
	src.Asset(2, t, false, CheckSortAsc([]int{1,2,5,4}))
}

func TestIsValid(t *testing.T) {
	bst := src.NewBSTree(1, 1)
	src.Asset(1, t, true, IsValid(bst.GetRoot()))
}

func TestCheckSortAsc2(t *testing.T) {
	a := []int{1, 0, 3}
	for _, v := range a {
		fmt.Println(v)
	}
}