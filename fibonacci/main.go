package main

import "fmt"

func main() {
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
func fibonacci(c chan<- int, quit <-chan int) {
	//定义x, y
	x, y := 0, 1
	for {
		select { //每次把x变量赋值给c chan
		case c <- x:
			x, y = y, x+y
		case <-quit: //收到退出信号, 就退出函数
			fmt.Println("quit")
			return
		}
	}
}
