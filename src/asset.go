package src

import "testing"

func TestAsset(t *testing.T) {
	tests := []struct{
		data []int //数据
		expect int //预期值
	}{
		{},
	}
	for index, item := range tests {
		if actual := Get(item.data); actual != item.expect {
			index ++
			t.Errorf("index:%d, expect:%d, actual:%d\n", index, item.expect, actual)
		}
	}
}
func Get(arr []int) int {
	return 1
}