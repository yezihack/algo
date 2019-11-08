package main

//综上双亲链与孩子链的优缺点,我们将其结合创建一链双亲孩子链结构

//defined single linked node
type PCNode struct {
	parentIndex int //parents index position
	next        *PCNode
}

//defined element
type PCBox struct {
	data      string  //node value
	parent    int     //parents index
	firstNode *PCNode //save linked node
}

//defined tree
type PCTree struct {
	nodes     []*PCBox //save node info
	rootIndex int      //record root node index
	count     int      //node number
}
