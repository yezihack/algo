package fibonacci

import "fmt"

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
func fibonacci2(want int) int {
	x, y := 0, 1
	i := 0
	for i < want {
		i++
		x, y = y, x+y
		fmt.Printf("i:%d, x:%d \n", i, x)
	}
	return x
}
