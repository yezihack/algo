package string_string

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"testing"
)

func TestBFSearch(t *testing.T) {
	fmt.Println(BFSearch("aabbcc", "abb"))
}

func TestBFSearch2(t *testing.T) {
	tests := []struct {
		mainStr  string
		childStr string
		expect   int
	}{
		{
			"aabbcc",
			"abb",
			1,
		},
		{
			"qwe",
			"q",
			0,
		},
		{
			"cat",
			"oo",
			-1,
		},
		{
			"pbd",
			"d",
			2,
		},
	}
	for _, tt := range tests {
		idx := BFSearch(tt.mainStr, tt.childStr)
		if idx != tt.expect {
			t.Errorf("expect:%d, actual:%d\n", tt.expect, idx)
		}
	}
}

func TestBFS(t *testing.T) {
	src.Asset(1, t, 2, Bfs("abcdefg", "cde"))
	src.Asset(2, t, 0, Bfs("abcdefg", "abc"))
	src.Asset(3, t, 5, Bfs("abcdefg", "fg"))
}