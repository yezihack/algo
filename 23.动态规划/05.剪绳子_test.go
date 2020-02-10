package _3_动态规划

import (
	"fmt"
	"testing"
)

func TestMaxCut(t *testing.T) {
	fmt.Println(MaxCut(8))
}

func TestTaiXinCut(t *testing.T) {
	fmt.Println(TaiXinCut(8))
	fmt.Println(TaiXinCut(9))
}
func TestDpCut(t *testing.T) {
	fmt.Println(DpCut(8))
	//fmt.Println(DpCut(9))
}
