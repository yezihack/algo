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
	fmt.Println(lst)
}
