package Recurive

import (
	"fmt"
	"testing"
)

//求fab数列
//1,1,2,3,5
//递推公式: f(n) = f(n-1) + f(n-2)
//2^n
func fib(n int) int {
	if n < 1 {
		return 0
	}
	if n < 2 { //n==1, n==2时只返回他本身. .
		return 1
	}
	return fib(n-1) + fib(n-2)
}
func TestFib2(t *testing.T) {
	fmt.Println(fib(7))
}
func fibMap(n int, cache map[int]int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if _, ok := cache[n]; !ok {
		cache[n] = fibMap(n-1, cache) + fibMap(n-2, cache)
	}
	return cache[n]
}
func TestFibMap(t *testing.T) {
	cache := make(map[int]int)
	fmt.Println(fibMap(7, cache))
}
func BenchmarkFibMap(b *testing.B) {
	cache := make(map[int]int)
	for i := 0; i < b.N; i++ {
		fibMap(i, cache)
	}
}
func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(i)
	}
}
//fib最快速度O(logn)
func power(a, n int) int {
	if n == 0 {
		return 1
	}
	r := power(a, n / 2)
	if (n & 1) == 0 {
		return r * r * a
	}
	return r * r
}
func TestPower(t *testing.T) {
	fmt.Println(power(1, 6))
}
//fib最快速度O(logn)
func power2(a, n int) int {
	if n == 0 {
		return 1
	}
	res, tmp := 1, a
	for n != 0 {
		if (n & 1) == 0 {
			res *= tmp
		}
		n >>= 1
		tmp *= tmp
	}
	return res
}
