package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci2(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{
			name:  "测试1",
			value: 3,
			want:  2,
		},
		{
			name:  "测试2",
			value: 10,
			want:  55,
		},
		{
			name:  "测试1",
			value: 7,
			want:  13,
		},
	}
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			val := fibonacci2(item.value)
			if val != item.want {
				t.Errorf("want:%d, return_value:%d, input_value:%d", item.want, val, item.value)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	c := make(chan int)    //定义一个通道,存放数据的
	quit := make(chan int) //定义一个退出信号
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //打印结果
		}
		quit <- 0 //发送退出信号
	}()
	fibonacci(c, quit)
}
