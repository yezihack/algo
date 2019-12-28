/*
 * @Author: 百里
 * @Date: 2019-12-28 17:21:46
 * @LastEditTime : 2019-12-28 17:26:13
 * @LastEditors  : 百里
 * @Description: 链表倒数第k个节点
 * @FilePath: \yezihack\algo\04.链表\06.链表倒数第k个节点.go
 * @https://github.com/yezihack
 */
package linked

import "github.com/yezihack/algo/00.src"

func GetLinkedK(head *src.LinkedNode, k int) *src.LinkedNode {
	n := 0
	fast := head
	for fast != nil {
		n ++
		fast = fast.Next
	}
	index := n - k
	slow := head
	for i := 0; i < index; i ++ {
		slow = slow.Next
	}
	return slow
}

//思路使用两个指针, 一个快指针,一个慢指针.
func GetLinkedKPoint(head *src.LinkedNode, k int) *src.LinkedNode {
	if head == nil {
		return nil
	}
	if k <= 0 {
		return nil
	}
	fast, slow := head.Next, head
	//先让fast先走k-1步
	n := 0
	for fast != nil && k-1 > n {
		n ++
		fast = fast.Next
	}
	//然后让fast和slow一起走,直到fast走到链表尾部.slow就是k位置
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
//第一个指针先走k-1步，然后第二个指针开始走
//当第一个指针移动到最后时，第二个指针正好指向倒数第k个结点，只需要遍历一遍链表，显然更高效。
func GetLinkedKCount(head *src.LinkedNode, k int) *src.LinkedNode {
	if head == nil {
		return nil
	}
	if k <= 0 {
		return nil
	}
	fast, slow := head.Next, head
	n := 0
	for fast != nil {
		//第一个指针先走k-1步，然后第二个指针开始走
		if k - 1 > n {
			fast = fast.Next
			n ++
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return slow
}