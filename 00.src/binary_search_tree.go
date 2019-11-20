package src

import (
	"code.qschou.com/golang/errors"
	"fmt"
	"math"
	"strconv"
)

//定义二叉搜索树结点结构
type TreeNode struct {
	Data int
	LChild *TreeNode
	RChild *TreeNode
	Left *TreeNode
	Right *TreeNode
	Val int
}
//定义一个二叉搜索树
type BSTree struct {
	root *TreeNode//根结点
}
//实例一个二叉搜索树.
func NewBSTree(values ...int) *BSTree {
	//空树也是二叉搜索树.
	if len(values) == 0 {
		return &BSTree{}
	}
	//创建一个根结点
	root := &TreeNode{Data:values[0]}
	//将根结点保存起来
	tree := &BSTree{
		root:root,
	}
	//从第二个元素开始创建左右子树.
	for i := 1;i < len(values); i ++ {
		if tree.AddNode(values[i], root) != nil {
			return nil
		}
	}
	return tree
}
//获取根结点
func (b *BSTree) GetRoot() *TreeNode {
	return b.root
}
//添加结点
func (b *BSTree) AddNode(value int, root *TreeNode) error {
	if root == nil {
		return errors.New("tree is null")
	}
	node := &TreeNode{Data:value}
	//当结点小于根结点则插入根结点的左孩子结点上.
	if node.Data < root.Data {
		if root.LChild == nil {
			root.LChild = node
		} else {
			return b.AddNode(value, root.LChild)
		}
	}
	//当结点大于根结点则插入根结点的右孩子结点上
	if node.Data >= root.Data {
		if root.RChild == nil {
			root.RChild = node
		} else {
			return b.AddNode(value, root.RChild)
		}
	}
	return nil
}
//先序遍历树. 空结点采用#代替
//先根再左后右
func (b *BSTree) PreOrderTree(node *TreeNode) string {
	str := ""
	if node == nil {
		return ""
	}
	str += strconv.Itoa(node.Data)
	if node.LChild == nil {
		str += "#"
	}
	str += b.PreOrderTree(node.LChild)
	if node.RChild == nil {
		str += "#"
	}
	str += b.PreOrderTree(node.RChild)
	return str
}

//中序遍历树. 空结点采用#代替
//先左再根后右
func (b *BSTree) InOrderTree(node *TreeNode) string {
	str := ""
	if node == nil {
		return ""
	}
	if node.LChild == nil {
		str += "#"
	}
	str += b.InOrderTree(node.LChild)
	str += strconv.Itoa(node.Data)
	if node.RChild == nil {
		str += "#"
	}
	str += b.InOrderTree(node.RChild)
	return str
}

//后序遍历树. 空结点采用#代替
//先左再右后根
func (b *BSTree) PostOrderTree(node *TreeNode) string {
	str := ""
	if node == nil {
		return ""
	}
	if node.LChild == nil {
		str += "#"
	}
	str += b.PostOrderTree(node.LChild)
	if node.RChild == nil {
		str += "#"
	}
	str += b.PostOrderTree(node.RChild)
	str += strconv.Itoa(node.Data)
	return str
}
//判断是否是一棵合并二叉搜索树
func (b *BSTree) IsValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	val := root.Data
	fmt.Printf("min:%d, max:%d, val:%d\n", min, max, val)
	if min != int(math.MaxInt32) && val <= min {
		return false
	}
	if max != int(math.MaxInt32) && val >= max {
		return false
	}
	if !b.IsValid(root.LChild, min, val) {
		return false
	}
	if !b.IsValid(root.RChild, val, max) {
		return false
	}
	return true
}