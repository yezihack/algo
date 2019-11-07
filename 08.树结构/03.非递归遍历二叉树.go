package main

import "fmt"

//参考视频:https://www.bilibili.com/video/av35339922
//中序遍历的非递归二叉树算法
//用栈的思想实现非递归遍历
//解题思路: 访问一个结点时,判断是否为空,不为空则存入一个栈里,然后继续访问
//左子树,直到为空,则返回输出双亲结点的值,然后继续访问右子树的结点,以此循环往复
func InOrcerTraverseByFor(t *BiNode) {
	var (
		p     *BiNode       //定义一个移动结点
		Q     *BiNode       //定义一个存储结点
		stack = InitStack() //实例一个栈空间
	)
	p = t
	//判断p的指针是否为空, 或 这个栈里是否还有元素.
	for p != nil || !stack.StackEmpty() {
		if p != nil {
			stack.Push(p) //存入当前结点
			p = p.lChild  //继续访问它的左子树
		} else {
			Q = stack.Pop()     //弹出一个栈元素
			fmt.Println(Q.data) //输出结点值
			p = Q.rChild        //访问它的右子树
		}
	}
}
