package main

//孩子兄弟表示法,双称二叉表示法,或二叉链表表示法
//一个结点有三个域, 第1个指向第一个孩子, 第2个存储结点值, 第3个指向兄弟结点.
//see: https://www.bilibili.com/video/av35340332

type ChildBrotherNode struct {
	data        string            //结点值
	firstChild  *ChildBrotherNode //指向第一个孩子结点
	nextSibling *ChildBrotherNode //指向兄弟结点.
}
