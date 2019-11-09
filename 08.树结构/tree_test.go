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

//哈夫曼树单元测试
func TestNewHuffmanTree(t *testing.T) {
	w := []int{5, 29, 7, 8, 14, 23, 3, 11}
	hf := NewHuffmanTree(w)
	hf.CreateTree()
	for _, item := range hf.nodes {
		fmt.Printf("ID:%d,weight:%d,parent:%d,lChild:%d, rChild:%d\n",
			item.ID, item.weight, item.parent, item.lChildIndex, item.rChildIndex)
	}
	//s1, s2 := hf.Select2MinNode(3)
	//fmt.Println(s1, s2)
}

//哈夫曼树单元测试
func TestNewHuffmanTreeCode(t *testing.T) {
	w := []int{40, 30, 15, 5, 4, 3, 3}
	hf := NewHuffmanTree(w)
	hf.CreateTree()
	for _, item := range hf.nodes {
		fmt.Printf("ID:%d,weight:%d,parent:%d,lChild:%d, rChild:%d\n",
			item.ID, item.weight, item.parent, item.lChildIndex, item.rChildIndex)
	}
	//print huffman code
	chars := []byte("ABCDEFG")
	codes := hf.CreateCode(chars).codeMap
	fmt.Println(codes)
}

//编码操作
func TestHuffmanTree_Encode(t *testing.T) {
	chars := []byte("ILOVEYU")
	w := []int{1, 2, 3, 4, 5, 6, 7}
	hf := NewHuffmanTree(w)
	str := "ILOVEYOU"
	hf.CreateTree().CreateCode(chars)
	s := hf.Encode(str)
	fmt.Println(s)
	//解码
	s = "011001110101101110001010"
	result := hf.Decode(s)
	fmt.Println(result)
}

func TestHuffmanTree_CodeTraverse(t *testing.T) {
	str := "abcdef"
	h := new(HuffmanTree)
	s := h.StrTraverse(str)
	fmt.Println(s)
}
