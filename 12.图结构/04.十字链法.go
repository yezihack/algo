package main

/*
十字链主要存储是有向图
将出度与入度都存在一个结点上,方便查找.
缺点:求结点的度困难
*/
//defined arc node struct
type ArcLinkedNode struct {
	TailVexID int //弧尾结点ID
	HeadVexID int //弧头结点ID
	HeadLink *ArcLinkedNode //头指针
	TailLink *ArcLinkedNode//尾指针
}

//定义顶点结点的结构
type VertexNode struct {
	Data rune //node data
	FirstInArcPtr *ArcLinkedNode //入度边的指针
	FirstOutArcPtr *ArcLinkedNode//出度边的指针
}
type AdJoinTable struct {
	Table []*VertexNode //邻接表切片
	VerNum int //vertex number
	ArcNum int //arc number
}
