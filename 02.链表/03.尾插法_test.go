package linked

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTailInsertLinked(t *testing.T) {
	arr := []int{1,2,3}
	head := TailInsertLinked(arr)
	PrintLinkedV2(head)
}

func PrintLinkedV2(head *linkedV2)  {
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