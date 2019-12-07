package _3_动态规划

import (
	"fmt"
	"testing"
)

func TestGetPrev(t *testing.T) {
	//打印做某任务还可以做前面的任务一种映射关系, 不过只取最近的一个可以做的任务
	prev := GetPrev()
	for _, item := range prev {
		fmt.Printf("taskId:%d, prevId:%d\n", item.taskId, item.prevTaskId)
	}
}
