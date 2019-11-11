package graph

import (
	"fmt"
	"testing"
)

func TestNewALGraph(t *testing.T) {
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
	//验证一下A, A顶点指向B和C
	vertexA := lst[0]
	expectA := rune('A')
	if vertexA.data != expectA {
		t.Errorf("expect:%c,actual:%c\n", expectA, vertexA.data)
	}
	//验证边
	nextPtr := vertexA.firstArc
	if nextPtr.vertexID != 3 {
		t.Errorf("expect:%d,actual:%d\n", 3, nextPtr.vertexID)
	}
	if nextPtr.nextArc.vertexID != 1 {
		t.Errorf("expect:%d,actual:%d\n", 1, nextPtr.nextArc.vertexID)
	}
}
