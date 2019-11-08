package tree

import (
	"testing"
)

func TestNewBiTree(t *testing.T) {
	//先序遍历算法
	preStr1 := "AB##C##"
	preT := NewBiTree(preStr1, PreOrderTraverse)
	preTree := preT.CreateTree()
	preStr2 := preT.Tree2String(preTree)
	if preStr1 != preStr2 {
		t.Errorf("expect:%s, actual:%s\n", preStr1, preStr2)
	}
	//中序遍历
	inStr1 := "#B#A#C#"
	inT := NewBiTree(inStr1, InOrderTraverse)
	inTree := inT.CreateTree()
	inStr2 := inT.Tree2String(inTree)
	if preStr1 != preStr2 {
		t.Errorf("expect:%s, actual:%s\n", inStr1, inStr2)
	}
	//后序遍历
	postStr1 := "#B#A#C#"
	postT := NewBiTree(inStr1, PostOrderTraverse)
	postTree := postT.CreateTree()
	postStr2 := postT.Tree2String(postTree)
	if preStr1 != preStr2 {
		t.Errorf("expect:%s, actual:%s\n", postStr1, postStr2)
	}
}

//先序,中序,后序不同的遍历方法验证
func TestBiTreeNode_Tree2String(t *testing.T) {
	// 创建一个根结点
	root := new(BiTreeNode)
	root.data = "A"
	//创建一个左子树
	ltree := new(BiTreeNode)
	ltree.data = "B"
	root.left = ltree
	//创建一个右子树
	rtree := new(BiTreeNode)
	rtree.data = "C"
	root.right = rtree
	//再创建一个左子树给做右子树的孩子
	ltree2 := new(BiTreeNode)
	ltree2.data = "D"
	rtree.left = ltree2

	/*************************先序遍历*******************************/
	preS := "AB##CD###"
	root.TraverseType = PreOrderTraverse
	treeString := root.Tree2String(root)
	if preS != treeString {
		t.Errorf("pre: expect:%s, actual:%s\n", preS, treeString)
	}
	/*************************中序遍历*******************************/
	inS := "#B#A#D#C#"
	root.TraverseType = InOrderTraverse
	treeStringIn := root.Tree2String(root)
	if inS != treeStringIn {
		t.Errorf("pre: expect:%s, actual:%s\n", inS, treeStringIn)
	}
	/*************************后序遍历*******************************/
	postS := "##B##D#CA"
	root.TraverseType = PostOrderTraverse
	treeStringPost := root.Tree2String(root)
	if postS != treeStringPost {
		t.Errorf("pre: expect:%s, actual:%s\n", postS, treeStringPost)
	}
}

//复制单元测试
func TestBiTreeNode_Copy(t *testing.T) {
	preStr := "AB##CD##"
	preT := NewBiTree(preStr, PreOrderTraverse)
	preTree := preT.CreateTree()
	copyTree := preT.Copy(preTree)
	t1 := preT.Tree2String(preTree)
	t2 := preT.Tree2String(copyTree)
	if t1 != t2 {
		t.Errorf("pre: expect:%s, actual:%s\n", t1, t2)
	}
}

//计算树的高度
func TestBiTreeNode_Depth(t *testing.T) {
	preStr := "AB##CD###"
	preT := NewBiTree(preStr, PreOrderTraverse)
	preTree := preT.CreateTree()
	depth := preT.Depth(preTree)
	if depth != 3 {
		t.Errorf("pre: expect:%d, actual:%d\n", 3, depth)
	}
}

//计算树的所有结点数
func TestBiTreeNode_NodeCount(t *testing.T) {
	preStr := "AB##CD##E##"
	preT := NewBiTree(preStr, PreOrderTraverse)
	preTree := preT.CreateTree()
	n := preT.NodeCount(preTree)
	if n != 5 {
		t.Errorf("pre: expect:%d, actual:%d\n", 5, n)
	}
}

//计算树的叶子结点数
func TestBiTreeNode_LeadCount(t *testing.T) {
	preStr := "AB##CD##E##"
	preT := NewBiTree(preStr, PreOrderTraverse)
	preTree := preT.CreateTree()
	n := preT.LeadCount(preTree)
	if n != 3 {
		t.Errorf("pre: expect:%d, actual:%d\n", 3, n)
	}
}
