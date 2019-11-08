package main

//利用数组存储树结构,parent指针,指向他双亲的位置
//优点:找双亲容易
//缺点:找孩子难
//see: https://www.bilibili.com/video/av35340136

//定义一个双亲结点
type ParentNode struct {
	data   string //结点值
	parent int    //双亲位置
}

//定义一棵树
type ParentTree struct {
	nodes      []*ParentNode //一个切片存储结点
	rootIndex  int           //根结点位置
	totalCount int           //结点个数
}

func NewParentTree() *ParentTree {
	return &ParentTree{
		rootIndex: -1,
	}
}
