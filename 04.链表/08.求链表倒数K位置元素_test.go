package linked

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

type minLink struct {
	val int
	next *minLink
}
//题目: 求链表倒数k位置的元素.
//解题思路: 快慢指针实现
//先让快指针走k个位置, 然后一起走.
// 当快指针走到头,则慢指针指向的位置就是所求的结果
//时间O(n), 空间O(n)
func GetLinkedByBackwards(k int, head *minLink) int {
	//判断边界值0
	if k <= 0 {
		log.Fatal("k less than or equal to zero")
		return -1
	}
	//判断为空的情况下
	if head == nil {
		log.Println("head is nil")
		return -1
	}
	//声明二个慢与快指针
	var slow, fast *minLink
	fast = head //先让快指针走k步
	for i := 0; i < k; i ++ {
		if fast == nil { //判断指针的长度小于k大小.则返回-1
			log.Println("k > linked length")
			return -1
		}
		//这里需要注意位置, 最后一个值时, next is nil
		fast = fast.next
	}
	slow = head //定义一下慢指针的初使位置.然后让其一起走.
	for fast != nil {//当快指针走到头,就停止循环.
		fast = fast.next //一起走
		slow = slow.next//一起走
	}
	return slow.val //返回指针值.
}
func TestGetLinkedByBackwards(t *testing.T) {
	l0 := &minLink{val:0}
	l1 := &minLink{val:1}
	l0.next = l1
	l2 := &minLink{val:2}
	l1.next = l2
	l3 := &minLink{val:3}
	l2.next = l3
	l4 := &minLink{val:4}
	l3.next = l4
	tests := []struct{
		index int
		testValue int //测试的值
		expect int //预期返回的值
	}{
		{1, 5, 0},
		{2,4, 1},
		{3,3, 2},
		{4,2, 3},
		{5,1, 4},
		{6,10, -1},
		{7,-10, -1},
	}
	for _, help := range tests {
		if val := GetLinkedByBackwards(help.testValue, l0); val != help.expect {
			t.Errorf("index:%d, expect:%d, actual:%d\n", help.index,help.expect, val)
		}
	}
}
//思路二: 先反转链表,再求K位置
func GetKIndexValue(k int, head *minLink) int {
	//判断边界值和空值
	if k <= 0 {
		log.Fatal("k less than or equal to zero")
		return -1
	}
	if head == nil {
		return -1
	}
	//先返回链表操作
	//声明一个空链表指针
	var newLinked *minLink
	for head != nil {
		//断开链表
		// curr存储后部分链表
		curr := head.next
		//进行断开操作, 将链表的前部分断开
		head.next = newLinked //将head的next指向新链表
		newLinked = head //将整个head的前部分赋值给新链表
		head = curr //产生一个新的head继续循环
	}
	printMinLinked(newLinked)
	//return 0
	for i := 0; i < k; i ++ {
		if newLinked == nil {
			break
		}
		fmt.Println(i, newLinked)
		if i+1 == k {
			break
		}
		newLinked = newLinked.next
	}

	return newLinked.val
}
func TestGetKIndexValue(t *testing.T) {
	l0 := &minLink{val:0}
	l1 := &minLink{val:1}
	l0.next = l1
	l2 := &minLink{val:2}
	l1.next = l2
	l3 := &minLink{val:3}
	l2.next = l3
	l4 := &minLink{val:4}
	l3.next = l4
	tests := []struct{
		testValue int //测试的值
		expect int //预期返回的值
	}{
		{5, 5},
		//{4, 4},
		//{3, 3},
		//{2, 2},
		//{1, 1},
		//{10, -1},
		//{-10, -1},
	}
	for _, help := range tests {
		if val := GetKIndexValue(help.testValue, l0); val != help.expect {
			t.Errorf("expect:%d, actual:%d\n", help.expect, val)
		}
	}
}


func printMinLinked(head *minLink) {
	str := ""
	for head != nil {
		str += strconv.Itoa(head.val)
		head = head.next
		if head != nil {
			str += "=>"
		}
	}
	fmt.Println(str)
}