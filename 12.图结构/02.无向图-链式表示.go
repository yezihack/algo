package main

import "fmt"

//顶点,边之间的权重 三者之间的关系
type VerArcWeight struct {
	startVer rune //开始顶点
	endVer   rune //结束顶点
	weight   int  //权重值
}

type VerTexWeight struct {
	vertexValue rune
	weight      int
}

//顶点与顶点之间的边的结构定义
type arcNode struct {
	vertexID int      //顶点位置ID
	nextArc  *arcNode //另一个顶点对应的结点
	weight   int      //权重
}

//顶点结构,头结点
type verTextHeader struct {
	ID       int      //顶点位置ID
	data     rune     //顶点域值
	firstArc *arcNode //第一个对应的边
}

//图的定义
type ALGraph struct {
	adjList    []*verTextHeader //邻接表,即一维数组
	verNum     int              //顶点数
	arcNun     int              //边数量
	ver2verRel []*VerArcWeight  //顶点与顶点之间的关系,内含权重值
}

//输入总顶点数和总边数
//建立顶点表
//创建邻接表
//vexNum顶点数, arcNum边数
func NewALGraph(verArr []*VerArcWeight) *ALGraph {
	return &ALGraph{
		ver2verRel: verArr,
	}
}

//get graph
func (g *ALGraph) GetGraph() []*verTextHeader {
	return g.adjList
}

//初始化无向网
func (g *ALGraph) CreateUDG() {
	verList := g.GetVerTexList(g.ver2verRel) //获取所有顶点信息
	g.verNum = len(verList)                  //顶点数量
	g.arcNun = len(g.ver2verRel)             //边的数量
	//初始所有化顶点
	for i := 0; i < g.verNum; i++ {
		ver := verTextHeader{
			ID:       i,                      //记录顶点位置ID
			data:     verList[i].vertexValue, //赋值顶点值
			firstArc: new(arcNode),           //初始化表头结点的指针域
		}
		g.adjList = append(g.adjList, &ver)
	}
	//构造所有顶点之间的邻接表
	for i := 0; i < g.arcNun; i++ {
		rel := g.ver2verRel[i]         //获取顶点与顶点之间的关系
		x := g.LocateVex(rel.startVer) //查找顶点的位置ID
		y := g.LocateVex(rel.endVer)   //查找顶点的位置ID

		p1 := new(arcNode)                 //初始化一个边
		p1.vertexID = y                    //设置顶点边对应的顶点ID
		p1.nextArc = g.adjList[i].firstArc //使用头插法,将上一个first边接到新的边后面
		g.adjList[i].firstArc = p1         //将新的边接入到顶点后面
		/**********无向图,有两个方向,如果是有向图,只设置上面或下面即可******************************/

		p2 := new(arcNode)                 //因为无向图,是双向的,需要再设置一下另一顶点的边
		p2.vertexID = x                    //设置顶点边对应的顶点ID
		p2.nextArc = g.adjList[y].firstArc //使用头插法,将上一个first边接到新的边后面
		g.adjList[y].firstArc = p2         //将新的边接入到顶点后面
	}
}
func (g *ALGraph) LocateVex(vertexValue rune) int {
	for i := 0; i < g.verNum; i++ {
		if g.adjList[i].data == vertexValue {
			return g.adjList[i].ID
		}
	}
	return -1
}

//洗数据.从关系数组里找到所有顶点
//返回顶点值与权重
func (g *ALGraph) GetVerTexList(verArr []*VerArcWeight) []*VerTexWeight {
	verList := make([]*VerTexWeight, 0)
	verMap := make(map[rune]struct{})
	for _, item := range verArr {
		if _, ok := verMap[item.startVer]; !ok {
			verMap[item.startVer] = struct{}{}
			verList = append(verList, &VerTexWeight{
				vertexValue: item.startVer,
				weight:      item.weight,
			})
		}
		if _, ok := verMap[item.endVer]; !ok {
			verMap[item.endVer] = struct{}{}
			verList = append(verList, &VerTexWeight{
				vertexValue: item.endVer,
				weight:      item.weight,
			})
		}
	}
	return verList
}

func main() {
	rel := make([]*VerArcWeight, 0)
	rel = append(rel, &VerArcWeight{'A', 'B', 1})
	rel = append(rel, &VerArcWeight{'B', 'D', 1})
	rel = append(rel, &VerArcWeight{'C', 'A', 1})
	rel = append(rel, &VerArcWeight{'D', 'C', 1})
	graph := NewALGraph(rel)
	graph.CreateUDG()
	lst := graph.GetGraph()
	for _, item := range lst {
		fmt.Printf("vertexID:%d, vertexData:%s, nextPtr:%v, nextPtr:%v\n",
			item.ID, string(item.data), item.firstArc, item.firstArc.nextArc)
	}
	//show A
	fmt.Println(string(lst[0].data), lst[0].firstArc.vertexID, lst[0].firstArc.nextArc.vertexID)
}
