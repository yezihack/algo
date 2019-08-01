package dichotomy

import (
	"fmt"
	"sort"
	"testing"
)

func TestDichotomy(t *testing.T) {
	fmt.Println("第1种")
	lst := []int{
		1, 2, 3, 4, 5, 6,
	}
	sort.Ints(lst)
	fmt.Println("排完序的切片:", lst)
	tests := []struct {
		name string
		val  int
		want int
	}{
		{
			name: "测试1",
			val:  3,
			want: 2,
		},
		{
			name: "测试2",
			val:  18,
			want: -1,
		},
		{
			name: "测试3",
			val:  6,
			want: 5,
		},
		{
			name: "测试4",
			val:  1,
			want: 0,
		},
		{
			name: "测试5",
			val:  -100,
			want: -1,
		},
	}
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			_val := dichotomy(item.val, lst)
			if _val != item.want {
				t.Errorf("want:%d, wantErr:%d", item.want, _val)
			}
		})
	}
}

func TestDichotomy2(t *testing.T) {
	fmt.Println("第2种")
	lst := []int{
		1, 2, 3, 4, 5, 6,
	}
	sort.Ints(lst)
	fmt.Println("排完序的切片:", lst)
	tests := []struct {
		name string
		val  int
		want int
	}{
		{
			name: "测试1",
			val:  3,
			want: 2,
		},
		{
			name: "测试2",
			val:  18,
			want: -1,
		},
		{
			name: "测试3",
			val:  6,
			want: 5,
		},
		{
			name: "测试4",
			val:  1,
			want: 0,
		},
		{
			name: "测试5",
			val:  -100,
			want: -1,
		},
	}
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			_val := dichotomy2(item.val, lst)
			if _val != item.want {
				t.Errorf("want:%d, wantErr:%d", item.want, _val)
			}
		})
	}
}

func TestDichotomy3(t *testing.T) {
	fmt.Println("第3种")
	lst := []int{
		1, 2, 3, 4, 5, 6,
	}
	sort.Ints(lst)
	fmt.Println("排完序的切片:", lst)
	tests := []struct {
		name string
		val  int
		want int
	}{
		{
			name: "测试1",
			val:  3,
			want: 2,
		},
		{
			name: "测试2",
			val:  18,
			want: -1,
		},
		{
			name: "测试3",
			val:  6,
			want: 5,
		},
		{
			name: "测试4",
			val:  1,
			want: 0,
		},
		{
			name: "测试5",
			val:  -100,
			want: -1,
		},
	}
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			_val := dichotomy3(item.val, lst)
			if _val != item.want {
				t.Errorf("want:%d, wantErr:%d", item.want, _val)
			}
		})
	}
}

func TestDichotomy4(t *testing.T) {
	fmt.Println("第4种")
	lst := []int{
		1, 2, 3, 4, 5, 6,
	}
	sort.Ints(lst)
	fmt.Println("排完序的切片:", lst)
	tests := []struct {
		name string
		val  int
		want int
	}{
		{
			name: "测试1",
			val:  3,
			want: 2,
		},
		{
			name: "测试2",
			val:  18,
			want: -1,
		},
		{
			name: "测试3",
			val:  6,
			want: 5,
		},
		{
			name: "测试4",
			val:  1,
			want: 0,
		},
		{
			name: "测试5",
			val:  -100,
			want: -1,
		},
	}
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			_val := Find(item.val, lst)
			if _val != item.want {
				t.Errorf("want:%d, wantErr:%d", item.want, _val)
			}
		})
	}
}
