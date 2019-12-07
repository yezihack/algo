package _3_动态规划

import (
	"testing"
)

func TestGetMaxValueGoodsList(t *testing.T) {
	g := []*Goods{
		&Goods{},
		&Goods{"水水", 3, 10},
		&Goods{"书籍", 1, 3},
		&Goods{"食物", 2, 9},
		&Goods{"夹克", 2, 5},
		&Goods{"相机", 1, 6},
	}

	GetMaxValueGoodsList(g, 6)
}