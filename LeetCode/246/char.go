package _46

import (
	"sort"
)
//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
type ByteSort []byte
func (b ByteSort) Len() int {
	return len(b)
}
func (b ByteSort) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByteSort) Less(i, j int) bool {
	return b[i] < b[j]
}

func IsAnagram(s string, t string) bool {
	sb := []byte(s)
	tb := []byte(t)
	sort.Sort(ByteSort(sb))
	sort.Sort(ByteSort(tb))
	if len(sb) != len(tb) {
		return false
	}
	i := 0
	for i < len(sb) {
		if sb[i] != tb[i] {
			return false
		}
		i ++
	}
	return true
}

func IsAnagramByMap(s, t string) bool {
	//不相等则直接返回false
	if len(s) != len(t) {
		return false
	}
	//26个字母大小数组
	var m [26]int
	for i := 0;i < len(s);i ++ {
		m[s[i] - 'a'] ++ //负责加
		m[t[i] - 'a'] -- //负责减
		//最终两个字符串相等则都归于0
	}
	//判断一下是否有不等于0的值.
	for _, num := range m {
		if num != 0 {
			return false
		}
	}
	return true
}