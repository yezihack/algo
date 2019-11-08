package main

//借用一个数组,数组里存储结点值
// 数组里再创建一个位置存储一个单链表,记录孩子结点的位置
//see:https://www.bilibili.com/video/av35340290
//优点:找孩子容易
//缺点:找双亲难

//定义一个孩子结点结构
type ChildNode struct {
	child int
	next  *ChildNode
}

//定义双亲结点结构
type ChildBox struct {
	data       string     //结点值域
	firstChild *ChildNode //指向一个单链表
}

//定义孩子链表
type ChildTree struct {
	nodes      []*ChildBox //存储双亲结点
	rootIndex  int         //记录根结点位置
	totalCount int         //记录结点个数.
}
