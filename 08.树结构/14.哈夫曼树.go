package main

import "fmt"

//什么是哈夫曼树,
//给定N个权值作为N个叶子结点，构造一棵二叉树，若该树的带权路径长度达到最小，
// 称这样的二叉树为最优二叉树，
// 也称为哈夫曼树(Huffman Tree)。
// 哈夫曼树是带权路径长度最短的树，权值较大的结点离根较近。
//see: https://www.bilibili.com/video/av35817749

//定义哈夫曼结点结构
type HuffmanNode struct {
	ID          int //位置
	weight      int //权重
	parent      int //双亲ID位置
	lChildIndex int //左孩子结点位置
	rChildIndex int //右孩子结点位置
}

//定义哈夫曼树
type HuffmanTree struct {
	nodes        []*HuffmanNode //存储所有结点
	nodesNumber  int            //构造后的结点数
	weightLength int            //存储权重值
}

//初始哈夫曼结构
// 需要注意,切片大小为2n-1
// 初始权重为n
func NewHuffmanTree(weight []int) *HuffmanTree {
	m := 2*len(weight) - 1 //构造后的结点数:2n-1
	nodes := make([]*HuffmanNode, m)
	for i := 0; i < len(weight); i++ {
		//huffman := new(HuffmanNode)
		//huffman.weight = weight[i]
		nodes[i] = &HuffmanNode{
			ID:     i,
			weight: weight[i],
		}
	}
	return &HuffmanTree{
		nodesNumber:  m,
		nodes:        nodes,
		weightLength: len(weight),
	}
}

//已经初始化了, 然后需要两两合并,只到剩余一个根结点
func (h *HuffmanTree) CreateTree() {
	for i := h.weightLength; i < h.nodesNumber; i++ {
		//创建一个新结点
		newNode := new(HuffmanNode)
		//将当前的i赋值给新结点的ID
		newNode.ID = i
		//找到最小的二个结点数.
		s1, s2 := h.Select2MinNode(i)
		fmt.Printf("minNode:s1-ID:%d,s1-weight:%d, s2-ID:%d,s2-weight:%d\n",
			s1.ID, s1.weight, s2.ID, s2.weight)
		//判断是否为空
		if s1 != nil {
			s1.parent = newNode.ID      //二个最小结点数中的其中一个结点,设置新结点ID
			newNode.lChildIndex = s1.ID //设置新结点的左孩子的ID
			newNode.weight = s1.weight  //设置新结点的权重值
		}
		if s2 != nil {
			s2.parent = newNode.ID      //二个最小结点数中的其中一个结点,设置合并后的结点ID
			newNode.rChildIndex = s2.ID //设置新结点的右孩子的ID
			newNode.weight += s2.weight //累加新结点的权重值
		}
		h.nodes[i] = newNode //将新结点加入切片中
	}
}

//找到两个最小的结点
//maxIndex:从0到最manIndex中间找最小值
func (h *HuffmanTree) Select2MinNode(maxIndex int) (s1 *HuffmanNode, s2 *HuffmanNode) {
	//先找一个s1的最小值
	s1 = new(HuffmanNode)
	s1.weight = 1<<32 - 1
	s2 = s1
	for i := 0; i < maxIndex; i++ {
		if h.nodes[i].parent > 0 { //已经有双亲结点的忽略
			continue
		}
		//找没有双亲的结点的最小值
		if s1.weight > h.nodes[i].weight {
			s2 = s1
			s1 = h.nodes[i]
		} else if s2.weight > h.nodes[i].weight {
			s2 = h.nodes[i]
		}
	}
	return
}
