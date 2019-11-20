package _8

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
)

//判断一个数组是否是升序.思路.后面比前面.
func CheckSortAsc(a []int) bool {
	for i := 0; i < len(a); i ++ {
		//防止越界
		if i < len(a) -1 && a[i] > a[i+1] {
			return false
		}
	}
	return true
}
var out []int
func InOrder(node *src.TreeNode) {
	if node == nil {
		return
	}
	InOrder(node.LChild)
	out = append(out, node.Data)
	InOrder(node.RChild)
}

func IsValid(node *src.TreeNode) bool {
	if node == nil {
		return true
	}
	//中序遍历
	out = make([]int, 0)
	InOrder(node)
	fmt.Println(out)
	return CheckSortAsc(out)
}