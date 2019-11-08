package main

import (
	"fmt"
	"testing"
)

func TestBNode_Tree(t *testing.T) {
	//先序排列的树结构
	s := "AB##C#D##"
	node := NewROOTNode(s)
	tree := node.createTree()
	str := node.TreeToString(tree)
	if s != str {
		t.Errorf("check tree. expect:equal, actual:%v\n", str)
	}
	if tree.data != "A" {
		t.Errorf("expect:%s, actual:%s\n", "A", tree.data)
	}
	if tree.left.data != "B" {
		t.Errorf("expect:%s, actual:%s\n", "B", tree.data)
	}
	if tree.right.data != "C" {
		t.Errorf("expect:%s, actual:%s\n", "C", tree.data)
	}
	if tree.right.right.data != "D" {
		t.Errorf("expect:%s, actual:%s\n", "D", tree.data)
	}
}
func TestBNode_Copy(t *testing.T) {
	//创建一棵二叉树
	root := NewROOTNode("AB##C##")
	tree1 := root.createTree()
	strTree1 := root.TreeToString(tree1)
	fmt.Println(strTree1)

	//复制二叉树
	tree2 := root.Copy(tree1)
	strTree2 := root.TreeToString(tree2)
	fmt.Println(strTree2)

	//判断是否是一棵树
	if strTree1 != strTree2 {
		t.Errorf("expect:%s, actual:%s\n", strTree1, strTree2)
	}

}
