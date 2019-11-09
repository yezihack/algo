package tree

import (
	"fmt"
	"testing"
)

func TestNewHuffmanTree(t *testing.T) {
	arr := []int{8, 2, 3, 9, 10}
	fmt.Println(arr)
	ht := NewHuffmanTree(arr)
	ht.CreateTree()
	for _, item := range ht.nodes {
		fmt.Printf("ID:%d,weight:%d,parent:%d,lChild:%v, rChild:%v\n",
			item.ID, item.weight, item.parent, item.lChild, item.rChild)
	}
}
func TestHuffmanTree_CreateCode(t *testing.T) {
	arr := []int{8, 2, 3, 9}
	chars := []byte{'A', 'B', 'C', 'D'}
	h := NewHuffmanTree(arr)
	h.CreateTree()
	codes := h.CreateCode(chars)
	fmt.Println(codes)
}
