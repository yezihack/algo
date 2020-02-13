/*
 * @Author: 百里
 * @Date: 2020-02-10 22:49:50
 * @LastEditTime : 2020-02-11 16:17:52
 * @LastEditors  : 百里
 * @Description:
 * @FilePath: \leetcode\02.链表\01.单链表.go
 * @https://github.com/yezihack
 */
package linked

type ISingleLinked interface {
	InsertHead()
	Append()
}

//定义链表节点
type SingleLinkedNode struct {
	Next *SingleLinkedNode //单链表只有一个后继节点
	Data int               //节点数据
}
type SingleLinked struct {
	head *SingleLinkedNode //只有一个头指针
}
