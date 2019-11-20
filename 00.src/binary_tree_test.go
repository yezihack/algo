package src

import (
	"fmt"
	"testing"
)

func TestCreateBTree(t *testing.T) {
	a := "108##12##"
	node := CreateBTree(a)
	fmt.Println(node.Data, node.LChild.Data)
	fmt.Println(node.Data, node.RChild.Data)
	fmt.Println(PrintBTree(node))
}
