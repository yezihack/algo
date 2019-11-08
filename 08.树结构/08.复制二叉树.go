package main

//定义一个结点结构
type bNode struct {
	data    string //结点值
	left    *bNode //左子树
	right   *bNode //右子树
	tCount  int    //读取字符串的指针
	strTree string //字符化的二叉树结构,空结点使用#代替
}

//实例化一个根结点
func NewROOTNode(strTree string) *bNode {
	return &bNode{
		strTree: strTree,
		tCount:  -1,
	}
}

//创建树
func (b *bNode) createTree() *bNode {
	var node *bNode
	b.tCount++
	if b.tCount >= len(b.strTree) {
		return nil
	}
	s := b.strTree[b.tCount]
	if s == '#' {
		return nil
	} else {
		node = new(bNode)
		node.data = string(s)
		node.left = b.createTree()
		node.right = b.createTree()
	}
	return node
}

//复制一棵树
func (b *bNode) Copy(t *bNode) *bNode {
	if t == nil {
		return nil
	}
	var node *bNode
	if t != nil {
		node = new(bNode)
		node.data = t.data
		node.left = b.Copy(t.left)
		node.right = b.Copy(t.right)
	}
	return node
}

//将树还原成字符串形式,无孩子结点使用#代替
func (b *bNode) TreeToString(t *bNode) string {
	if t == nil {
		return ""
	}
	var str string
	//fmt.Print(t.data)
	str += t.data
	if t.left == nil {
		str += "#"
	}
	if t.right == nil {
		str += "#"
	}
	str += b.TreeToString(t.left)
	str += b.TreeToString(t.right)
	return str
}
