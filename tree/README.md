## 树结构

## 目录
1. 二叉树的操作
    1. 二叉树的创建
    1. 还原二叉树
    1. 复制二叉树
    1. 计算树的高度
    1. 计算树的所有结点数
    1. 计算树的叶子结点数
1. 哈夫曼树
    1. 创建哈夫曼树
    1. 找到两个最小的权重的结点
    1. 哈夫曼编码与解码

### [实现二叉树的基本操作](binary_tree.go)
> 参考: https://www.bilibili.com/video/av35340048
1. 创建一棵树 CreateTree
```
func (t *BiTreeNode) CreateTree() *BiTreeNode {
	t.iCount++
	if t.iCount >= len(t.strTree) {
		return nil
	}
	var node *BiTreeNode
	s := t.strTree[t.iCount]
	if s == '#' {
		return nil
	} else {
		node = new(BiTreeNode)
        node.left = t.CreateTree()
        node.right = t.CreateTree()
        node.data = string(s)
	}
	return node
}
```
2. 还原一棵树 Tree2String
```
func (t *BiTreeNode) Tree2String(tree *BiTreeNode) string {
	if tree == nil {
		return ""
	}
	var str string
	str += tree.data
	if tree.left == nil {
        str += "#"
    }
    str += t.Tree2String(tree.left)
    if tree.right == nil {
        str += "#"
    }
    str += t.Tree2String(tree.right)    
	return str
}
```
3. 复制一棵树 Copy
```
func (t *BiTreeNode) Copy(tree *BiTreeNode) *BiTreeNode {
	if tree == nil {
		return nil
	}
	node := new(BiTreeNode)
	node.data = tree.data
	if tree.left != nil {
		node.left = t.Copy(tree.left)
	}
	if tree.right != nil {
		node.right = t.Copy(tree.right)
	}
	return node
}
```
4. 计算树的高度 Depth
```
func (t *BiTreeNode) Depth(tree *BiTreeNode) int {
	if tree == nil {
		return 0
	}
	m := t.Depth(tree.left)
	n := t.Depth(tree.right)
	if m > n {
		return m + 1
	}
	return n + 1
}
```
5. 计算树的所有结点数 NodeCount
```
func (t *BiTreeNode) NodeCount(tree *BiTreeNode) int {
	if tree == nil {
		return 0
	}
	m := t.NodeCount(tree.left)
	n := t.NodeCount(tree.right)
	return m + n + 1
}
```
6. 计算树的叶子结点数 LeadCount
```
func (t *BiTreeNode) LeadCount(tree *BiTreeNode) int {
	if tree == nil {
		return 0
	}
	if tree.left == nil && tree.right == nil {
		return 1
	}
	return t.LeadCount(tree.left) + t.LeadCount(tree.right)
}
```

### [哈夫曼树实现](huffman_tree.go)
> 哈夫曼是无损的编码与解码.属于前缀编码及最优产缀编码.
> see:https://www.bilibili.com/video/av35817749
1. 核心代码,创建哈夫曼树,两两合并
```
//已经初始化了, 然后需要两两合并,只到剩余一个根结点
func (h *HuffmanTree) CreateTree() {
	for i := h.weightLength; i < h.nodesNumber; i++ {
		//创建一个新结点
		newNode := new(HuffmanNode)
		//将当前的i赋值给新结点的ID
		newNode.ID = i
		//找到最小的二个结点数.
		s1, s2 := h.Find2MinNode(i)
		fmt.Printf("minNode:s1-ID:%d,s1-weight:%d, s2-ID:%d,s2-weight:%d\n",
			s1.ID, s1.weight, s2.ID, s2.weight)
		newNode.weight = s1.weight + s2.weight //新结点的权重就是两个孩子权重之和
		newNode.lChild = s1                    //新结点的左孩子
		newNode.rChild = s2                    //新结点的右孩子
		s1.parent = newNode.ID                 //设置左孩子的双亲ID
		s2.parent = newNode.ID                 //设置右孩子的双亲ID
		h.nodes[i] = newNode                   //将新结点加入切片中
	}
}
``` 
2. 找到两个最小的权重代码
> 1个循环实现
```
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
```
3. 哈夫曼编码与解码
```
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
```