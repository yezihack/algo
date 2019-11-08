package main

//线索二叉线.
//主要是利用叶子结点的空指针标记前趋与后续
/*
二叉树链表
具有n个结点的二叉链表中, 一共有2n个指针域.因为每一个结点都有二个指针域
一个二叉链表中,只有n-1个孩子 ,为什么减1,因为根结点不是孩子.
n-1用于表示结点的左右孩子数, 其余n+1个指针域为空
*/
//参考: https://www.bilibili.com/video/av35340088
type BiThreadTreeNode struct {
	data   string            //结点值
	lChild *BiThreadTreeNode //左孩子
	rChild *BiThreadTreeNode //右孩子
	lTag   int               //0表示有左孩子,1表示前趋
	rTag   int               //0表示有右孩子,1表示后续
}
