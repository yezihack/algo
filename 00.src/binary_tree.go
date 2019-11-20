package src

import "fmt"

//二叉树定义

type BTreeNode struct {
	Data interface{}  //node data value
	LChild *BTreeNode //left child
	RChild *BTreeNode //right child
}

func NewBTreeNode(data interface{}) *BTreeNode {
	return &BTreeNode{
		Data:data,
	}
}
var i = -1
func CreateBTree(arr string) *BTreeNode {
	i = i + 1
	if i >= len(arr) {
		return nil
	}
	var node *BTreeNode
	if arr[i] == '#' {
		return nil
	} else {
		node = new(BTreeNode)
		node.Data = string(arr[i])
		node.LChild = CreateBTree(arr)
		node.RChild = CreateBTree(arr)
	}
	return node
}
//打印二叉树.
func PrintBTree(root *BTreeNode) string {
	str := ""
	if root == nil {
		return ""
	}
	str += fmt.Sprint(root.Data)
	str += PrintBTree(root.LChild)
	if root.LChild == nil {
		str += "#"
	}
	str += PrintBTree(root.RChild)
	if root.RChild == nil {
		str += "#"
	}
	return str
}