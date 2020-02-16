/*
 * @Author: 百里
 * @Date: 2020-02-11 16:18:27
 * @LastEditTime: 2020-02-11 16:18:39
 * @LastEditors: 百里
 * @Description:
 * @FilePath: \leetcode\02.链表\02.头插法.go
 * @https://github.com/yezihack
 */
package linked

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHeadInsertLinked(t *testing.T) {
	arr := []int{1,2,3}
	head := HeadInsertLinked(arr)
	PrintLinked(head)
}
func PrintLinked(head *linked)  {
	s := ""
	for head != nil {
		s += strconv.Itoa(head.data)
		head = head.next
		if head != nil {
			s += "=>"
		}
	}
	fmt.Println(s)
}