package tree

import (
	"strings"
)

//什么是哈夫曼树,
//给定N个权值作为N个叶子结点，构造一棵二叉树，若该树的带权路径长度达到最小，
// 称这样的二叉树为最优二叉树，
// 也称为哈夫曼树(Huffman Tree)。
// 哈夫曼树是带权路径长度最短的树，权值较大的结点离根较近。
//see: https://www.bilibili.com/video/av35817749

//定义哈夫曼结点结构
type HuffmanNode struct {
	ID     int //位置
	weight int //权重
	parent int //双亲ID位置
	lChild int //左孩子结点位置
	rChild int //右孩子结点位置
}

//定义哈夫曼树
type HuffmanTree struct {
	nodes        []*HuffmanNode  //存储所有结点
	nodesNumber  int             //构造后的结点数
	weightLength int             //存储权重值
	code2CharMap map[string]byte //编码=>字符
	char2CodeMap map[byte]string //字符=>编码
}

//初始哈夫曼结构
// 需要注意,切片大小为2n-1
// 初始权重为n
func NewHuffmanTree(weight []int) *HuffmanTree {
	max := 2*len(weight) - 1 //构造后的结点数:2n-1
	nodes := make([]*HuffmanNode, max)
	for i := 0; i < len(weight); i++ {
		//huffman := new(HuffmanNode)
		//huffman.weight = weight[i]
		nodes[i] = &HuffmanNode{
			ID:     i,
			weight: weight[i],
			lChild: -1,
			rChild: -1,
		}
	}
	return &HuffmanTree{
		nodesNumber:  max,
		nodes:        nodes,
		weightLength: len(weight),
	}
}

//已经初始化了, 然后需要两两合并,只到剩余一个根结点
func (h *HuffmanTree) CreateTree() *HuffmanTree {
	for i := h.weightLength; i < h.nodesNumber; i++ {
		//创建一个新结点
		newNode := new(HuffmanNode)
		//将当前的i赋值给新结点的ID
		newNode.ID = i
		//找到最小的二个结点数.
		s1, s2 := h.Find2MinNode(i)
		//fmt.Printf("minNode:s1-ID:%d,s1-weight:%d, s2-ID:%d,s2-weight:%d\n",
		//	s1.ID, s1.weight, s2.ID, s2.weight)
		newNode.weight = s1.weight + s2.weight //新结点的权重就是两个孩子权重之和
		newNode.lChild = s1.ID                 //新结点的左孩子
		newNode.rChild = s2.ID                 //新结点的右孩子
		s1.parent = newNode.ID                 //设置左孩子的双亲ID
		s2.parent = newNode.ID                 //设置右孩子的双亲ID
		h.nodes[i] = newNode                   //将新结点加入切片中
	}
	return h
}

//生成哈夫曼对应的编码
//思路:从根结点开始找,左孩子则输出0, 右孩子则输出1,然后反转字符串得到huffman编码
func (h *HuffmanTree) CreateCode(chars []byte) *HuffmanTree {
	codes := make(map[byte]string)
	//只循环叶子结点数的次数
	for i := 0; i < h.weightLength; i++ {
		parent := h.nodes[i].parent //找到当前叶子结点的parent
		currID := i                 //记录一下当前ID
		code := ""                  //存储huffman code
		for parent != 0 {           //循环直到找到parent == 0即是根结点
			prevNode := h.nodes[parent]    //找到上一个结点
			if prevNode.lChild == currID { //判断是否是左孩子,是则标记为0
				code += "0"
			} else if prevNode.rChild == currID { //判断是否是右孩子,是则标记为1
				code += "1"
			}
			parent = prevNode.parent //继续找双亲结点
			currID = prevNode.ID     //重置当前的双亲结点ID
		}
		codes[chars[i]] = h.StrTraverser(code) //将得到的huffman code反转过来
		//我们是从叶子结点开始找的,我们正常编码是从根开始数的,所以需要反转一下
	}
	h.char2CodeMap = codes
	h.code2CharMap = h.MapTraverse(codes)
	return h
}

//编码
func (h *HuffmanTree) Encode(str string) string {
	var b strings.Builder
	for i := 0; i < len(str); i++ {
		s := byte(str[i])
		if code, ok := h.char2CodeMap[s]; ok {
			b.WriteString(code)
		}
	}
	return b.String()
}

//解码
func (h *HuffmanTree) Decode(codes string) string {
	var b strings.Builder
	i, j := 0, 1
	for j <= len(codes) {
		c := codes[i:j]
		if char, ok := h.code2CharMap[c]; ok {
			b.WriteByte(char)
			i = j
		} else {
			j++
		}
	}
	return b.String()
}

//字符串反转操作
func (h *HuffmanTree) StrTraverser(str string) string {
	l := len(str)
	b := []byte(str)
	for i := 0; i < l/2; i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
	return string(b)
}

//map转换
func (h *HuffmanTree) MapTraverse(m map[byte]string) map[string]byte {
	codeMap := make(map[string]byte)
	for char, code := range m {
		codeMap[code] = char
	}
	return codeMap
}

//找到两个最小的结点
//maxIndex:从0到最manIndex中间找最小值
func (h *HuffmanTree) Find2MinNode(maxIndex int) (s1 *HuffmanNode, s2 *HuffmanNode) {
	//先找一个s1的最小值
	s1 = new(HuffmanNode)
	s1.weight = 1<<32 - 1
	s2 = new(HuffmanNode)
	s2.weight = 1<<32 - 1
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
