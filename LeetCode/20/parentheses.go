package _0

import (
	"github.com/yezihack/algo/00.src"
)

func IsValid(s string) bool {
	//设置括号MAP,注意定义的顺序.如()右半边做为键,左半边做为值.有秒用.
	bm := map[byte]byte{
		')':'(',
		'}':'{',
		']':'[',
	}
	//申请一个栈
	stack := src.NewStack()
	//将字符串转化为byte
	sb := []byte(s)
	//循环sb字节切片
	for _, b := range sb {
		if _, ok := bm[b]; !ok { //判断是否在map里,不在则压栈
			stack.Push(b)
		} else {
			if val, ok := bm[b]; ok {//不存在map键里,则向栈中弹出一个元素进行比较.
				sVal := stack.Pop()//弹出栈
				if sVal != nil && sVal.(byte) != val { //防止栈里无元素.进行比较,如果不相等则直接false
					return false
				}
			}
		}
	}
	if stack.StackEmpty() {
		return true
	}
	return false
}
func IsValidBySlice(s string) bool {
	bm := map[byte]byte{
		')':'(',
		'}':'{',
		']':'[',
	}
	i := 0 //声明一个指针,使用跟踪stack切片的,模拟栈的操作.
	stack := make([]byte, len(s))
	sb := []byte(s)
	for _, b := range sb {
		if _, ok := bm[b]; !ok {
			stack[i] = b
			i ++
		} else {
			if val, ok := bm[b]; ok {
				if i != 0 {
					i --
				}
				sVal := stack[i]
				if sVal != val {
					return false
				}
			}
		}
	}
	if i == 0 {
		return true
	}
	return false
}
