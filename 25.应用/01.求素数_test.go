package _5_应用

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"math"
	"testing"
)

//什么是素数,也叫质数.
//大于1的自然数, 只能被1或自身整除的数才是素数.

//判断是否是素数
func IsPrime(i int) bool {
	for j := 2; j <= i; j ++ {
		if i % j == 0 && i != j {
			return false
		}
	}
	return true
}
func TestPrime(t *testing.T) {
	src.Asset(1, t, true, IsPrime(3))
	src.Asset(2, t, true, IsPrime(5))
	src.Asset(3, t, true, IsPrime(7))
	src.Asset(4, t, true, IsPrime(11))
	src.Asset(5, t, true, IsPrime(13))
	src.Asset(6, t, false, IsPrime(16))
}
//进一步改进 O(n/2) 相比之前方法提升一倍
func IsPrime2(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return false
	}
	for i := 2; i < n; i += 2 {
		fmt.Printf("i:%d\n", i)
		if n % i == 0 {
			return false
		}
	}
	return true
}
func TestPrime2(t *testing.T) {
	src.Asset(1, t, true, IsPrime2(3))
	src.Asset(2, t, true, IsPrime2(5))
	src.Asset(3, t, true, IsPrime2(7))
	src.Asset(4, t, true, IsPrime2(11))
	src.Asset(5, t, true, IsPrime2(13))
	src.Asset(6, t, false, IsPrime2(16))
}
//进一步改进 O(n/2) 相比之前方法提升一倍
func IsPrime3(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return false
	}
	for i := 2; i*i < n; i += 2 {
		fmt.Printf("i:%d\n", i)
		if n % i == 0 {
			return false
		}
	}
	return true
}
func TestPrime3(t *testing.T) {
	fn := func(n int) bool {
		return IsPrime3(n)
	}
	src.Asset(1, t, true, fn(3))
	src.Asset(2, t, true, fn(5))
	src.Asset(3, t, true, fn(7))
	src.Asset(4, t, true, fn(11))
	src.Asset(5, t, true, fn(13))
	src.Asset(6, t, false, fn(16))
}

//奇偶数里找. 还可以继续优化. 只在奇数里找.
func MakePrime(num int) []int {
	primes := make([]int, 0)
	var i, j int
	for i = 2; i < num; i ++{
		fmt.Printf("i:%d, j:%d\n", i, j)
		for j = 2; j <= i - 1; j ++{
			if i % j == 0 {
				break
			}
		}
		if j == i {
			primes = append(primes, i)
		}
	}
	return primes
}
func TestMakePrime(t *testing.T) {
	fmt.Println(MakePrime(10))
}
func MakePrime2(num int) []int {
	primes := make([]int, 0)
	var i, j int
	for i = 2; i < num; i ++{
		fmt.Printf("i:%d, j:%d\n", i, i/2)
		for j = 2; j <= i/2; j ++{ //过滤掉偶数.
			if i % j == 0 {
				break
			}
		}
		if j > i/2 {
			primes = append(primes, i)
		}
	}
	return primes
}
func TestMakePrime2(t *testing.T) {
	fmt.Println(MakePrime2(10))
}
func MakePrime3(num int) []int {
	primes := make([]int, 0)
	var i, j int
	for i = 2; i < num; i ++{
		fmt.Printf("i:%d, j:%d\n", i, i/2)
		for j = 2; j <= int(math.Sqrt(float64(i))); j ++{ //过滤掉偶数.
			if i % j == 0 {
				break
			}
		}
		if j > int(math.Sqrt(float64(i))) {
			primes = append(primes, i)
		}
	}
	return primes
}
func TestMakePrime3(t *testing.T) {
	fmt.Println(MakePrime3(10))
}