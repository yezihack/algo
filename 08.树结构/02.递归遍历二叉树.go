package main

import "fmt"

//先序遍历算法
func PreOrderTraverse(t *BiNode) {
	if t == nil {
		fmt.Println("tree is null")
		return
	} else {
		fmt.Println(t.data)
		PreOrderTraverse(t.lChild)
		PreOrderTraverse(t.rChild)
	}
}

//中序遍历算法
func InOrderTraverse(t *BiNode) {
	if t == nil {
		fmt.Println("tree is null")
		return
	} else {
		InOrderTraverse(t.lChild) //递归左子树
		fmt.Println(t.data)       //再访问根结点
		InOrderTraverse(t.rChild) //递归右子树
	}
}

//后序遍历算法
func PostOrderTraverse(t *BiNode) {
	if t == nil {
		fmt.Println("tree is nul")
		return
	}
	PostOrderTraverse(t.lChild) //递归左子树
	PostOrderTraverse(t.rChild) //递归右子树
	fmt.Println(t.data)         //访问根结点
}
