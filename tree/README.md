## 树结构

### 实现二叉树的基本操作
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

> 参考: https://www.bilibili.com/video/av35340048