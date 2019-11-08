package main

type ElementType interface{}

//二叉链表的结点定义
type BiNode struct {
	data   ElementType //存储数据
	lChild *BiNode     //左孩子结点
	rChild *BiNode     //右孩子结点
}

func NewBiNode(data ElementType) *BiNode {
	return &BiNode{
		data: data,
	}
}

//三叉链表的结点定义
type TriTNode struct {
	data   ElementType //存储数据
	lChild *BiNode     //左孩子结点
	rChild *BiNode     //右孩子结点
	parent *BiNode     //指向双亲结点
}
