package main

import "fmt"

type VerTextType byte //顶点类型
type ArcType int      //边的权值类型

//Adjacency Matrix Graph
type AMGraph struct {
	VerTexType []VerTextType //顶点表
	ArcType    [][]ArcType   //邻接矩阵
	verNum     int           //图的顶点数
	arcNum     int           //图的边数
}

func CreateUDN(verList []VerTextType) *AMGraph {
	max := len(verList)
	arcList := make([][]ArcType, max)
	for i := 0; i < max; i++ {
		arcCol := make([]ArcType, max)
		for j := 0; j < max; j++ {
			arcCol[j] = 1<<32 - 1
		}
		arcList[i] = arcCol
	}
	return &AMGraph{
		VerTexType: verList,
		ArcType:    arcList,
		verNum:     max,
	}
}
func main() {
	ver := []VerTextType{'A', 'B', 'C', 'D'}
	udn := CreateUDN(ver)
	fmt.Println(udn.ArcType)
}
