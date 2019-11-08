package tree

/*
实现一个二叉树的基本操作
1.创建一棵树 CreateTree
2.还原一棵树 Tree2String
3.复制一棵树 Copy
4.计算树的高度 Depth
5.计算树的所有结点数 NodeCount
6.计算树的叶子结点数 LeadCount

参考: https://www.bilibili.com/video/av35340048
*/

const (
	PreOrderTraverse  = iota //先序遍历
	InOrderTraverse          //中序遍历
	PostOrderTraverse        //后序遍历
)

//定义树结点
type BiTreeNode struct {
	data         string      //结点值
	left         *BiTreeNode //左子树
	right        *BiTreeNode //右子树
	iCount       int         //计算
	strTree      string      //字符串形式的树结构,如AB##C##, #代表空
	TraverseType int         //遍历类型
}

func NewBiTree(strTree string, traverseType int) *BiTreeNode {
	return &BiTreeNode{
		iCount:       -1,
		strTree:      strTree,
		TraverseType: traverseType,
	}
}

//创建二叉树
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
		switch t.TraverseType {
		case PreOrderTraverse:
			node.data = string(s)
			node.left = t.CreateTree()
			node.right = t.CreateTree()
		case InOrderTraverse:
			node.left = t.CreateTree()
			node.data = string(s)
			node.right = t.CreateTree()
		case PostOrderTraverse:
			node.left = t.CreateTree()
			node.right = t.CreateTree()
			node.data = string(s)
		}
	}
	return node
}

//将链树还原成字符串格式的树结构
func (t *BiTreeNode) Tree2String(tree *BiTreeNode) string {
	if tree == nil {
		return ""
	}
	var str string
	lPound := func() {
		if tree.left == nil {
			str += "#"
		}
	}
	rPound := func() {
		if tree.right == nil {
			str += "#"
		}
	}
	switch t.TraverseType {
	case PreOrderTraverse:
		str += tree.data
		lPound()
		str += t.Tree2String(tree.left)
		rPound()
		str += t.Tree2String(tree.right)
	case InOrderTraverse:
		str += t.Tree2String(tree.left)
		lPound()
		str += tree.data
		str += t.Tree2String(tree.right)
		rPound()
	case PostOrderTraverse:
		lPound()
		str += t.Tree2String(tree.left)
		str += t.Tree2String(tree.right)
		rPound()
		str += tree.data
	}
	return str
}

//复制一棵树
//思路:先将左子树复制,再复制右子树,递归实现
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

//计算树的高度
//分别计算左子树的高度与右子树的高度,两者进行比较,得最大值的那个
//然后再+1,因为根结点高度为1
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

//计算树的所有结点数
//先计算左子树的结点数,再计算右子树的结点数,然后相加,最后加上根结点1
func (t *BiTreeNode) NodeCount(tree *BiTreeNode) int {
	if tree == nil {
		return 0
	}
	m := t.NodeCount(tree.left)
	n := t.NodeCount(tree.right)
	return m + n + 1
}

//计算树的叶子结点数
//当左子树为空并右子树也为空即是叶子结点,否则不是.
//将左子树与右子树计算的叶子结点相加得出最终的叶子结点数.
func (t *BiTreeNode) LeadCount(tree *BiTreeNode) int {
	if tree == nil {
		return 0
	}
	if tree.left == nil && tree.right == nil {
		return 1
	}
	return t.LeadCount(tree.left) + t.LeadCount(tree.right)
}
