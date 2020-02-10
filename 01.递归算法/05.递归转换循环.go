package Recurive

//使用尾递归实现 fibonacci
func Fib(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

//使用循环实现 fibonacci
func FibByFor(n int) int {
	if n < 0 {
		return 0
	}
	s1, s2 := 0, 1
	for i := 3; i <= n; i++ {
		s1, s2 = s2, s1+s2
	}
	return s2
}
