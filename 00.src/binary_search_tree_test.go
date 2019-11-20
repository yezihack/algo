package src

import (
	"math"
	"testing"
)

func TestCreateBSTree(t *testing.T) {
	bst := NewBSTree(10, 8, 12)
	root := bst.GetRoot()
	Asset(1, t, 10, root.Data)
	Asset(2, t, 8, root.LChild.Data)
	Asset(3, t, 12, root.RChild.Data)

	Asset(4, t, nil, bst.AddNode(7, root))
	Asset(5, t, "1087###12##", bst.PreOrderTree(root))
	Asset(6, t, "#7#8#10#12#", bst.InOrderTree(root))
	Asset(7, t, "##7#8##1210", bst.PostOrderTree(root))

	Asset(8, t, true, bst.IsValid(root, int(math.MaxInt32), int(math.MaxInt32)))
}
