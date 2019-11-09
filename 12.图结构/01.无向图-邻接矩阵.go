package main

type VerTextType byte //顶点类型
type ArcType int      //边的权值类型

//Adjacency Matrix Graph
type AMGraph struct {
	VerTexType []VerTextType //顶点表
	ArcType    [][]ArcType   //邻接矩阵
	verNum     int           //图的当前点数
	arcNum     int           //图的边数
}
