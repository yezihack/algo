package main

import "fmt"

//使用尾递归实现 fibonacci
func fib(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

//使用循环实现 fibonacci
func fibByFor(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	s1, s2 := 1, 1
	for i := 3; i <= n; i++ {
		s1, s2 = s2, s1+s2
	}
	return s2
}
func main() {
	n := 7
	fmt.Println(fib(n))
	fmt.Println(fibByFor(n))
}
