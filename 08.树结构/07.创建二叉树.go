package main

import "fmt"

//参考: https://blog.csdn.net/qq_36183935/article/details/80259526
type TreeNode struct {
	data  string
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(data string) *TreeNode {
	return &TreeNode{
		data: data,
	}
}

//遍历二叉树,三种遍历法, 先序,中序,后序
func TraverseTree(id int, root *TreeNode) {
	if root == nil {
		return
	}
	switch id {
	case 1: //先序遍历
		print(root.data, ",")
		TraverseTree(id, root.left)
		TraverseTree(id, root.right)
	case 2: //中序遍历
		print(root.data, ",")
		TraverseTree(id, root.left)
		TraverseTree(id, root.right)
	case 3: //后序遍历
		print(root.data, ",")
		TraverseTree(id, root.left)
		TraverseTree(id, root.right)
	}
}

var i = -1

//创建一棵二叉树
func CreateBiTree(arr []string) *TreeNode {
	i = i + 1
	if i >= len(arr) {
		return nil
	}
	var node *TreeNode
	if arr[i] == "#" {
		return nil
	} else {
		node = NewTreeNode(arr[i])
		node.left = CreateBiTree(arr)
		node.right = CreateBiTree(arr)
	}
	return node
}
func main() {
	var treeIds = []string{"A", "B", "#", "#", "C", "#", "D"}
	t := CreateBiTree(treeIds)
	TraverseTree(1, t)
	fmt.Println()
	fmt.Println(t.data)
	fmt.Println(t.left.data)
	fmt.Println(t.right.data)
	fmt.Println(t.right.right.data)

}
