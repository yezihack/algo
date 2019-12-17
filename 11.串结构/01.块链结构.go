/*
利用块链存储,提高存储密度
存储密度公式=串值所占的存储/实际分配的存储
*/
package string_string


import "unsafe"

//定义一个块链结点
type ChunkNode struct {
	ch   []byte
	next *ChunkNode
}

func NewChunkNode(data string) *ChunkNode {
	c := &ChunkNode{
		ch: make([]byte, 0),
	}
	c.ch = append(c.ch, string2byte(data)...)
	return c
}
func byte2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func string2byte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

type ChunkLinked struct {
	Head *ChunkNode
	Tail *ChunkNode
}
