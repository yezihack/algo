package main

//二叉排序树的递归查找

type BinaryNode struct {
	Key int//键
	Data int //值
	LChild *BinaryNode//左孩子
	RChild *BinaryNode //右孩子
}
//递归实现二叉排序树.
func SearchBST(tree *BinaryNode, data int) *BinaryNode {
	if tree == nil {
		return nil
	} else if data < tree.Data {
		return SearchBST(tree.LChild, data)
	} else if data > tree.Data {
		return SearchBST(tree.RChild, data)
	} else {
		return tree
	}
}
