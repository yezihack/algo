package main

import "testing"

func TestBFSearch(t *testing.T) {
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
