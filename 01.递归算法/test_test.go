package Recurive

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestQSort(t *testing.T) {
	a := []int{4, 2, 7, 9, 1, 8, 5, 6, 3}
	fmt.Println("sort:", a)
	QSort(a, 0, len(a)-1)
	fmt.Println("sorted", a)
}
func TestQuickSort(t *testing.T) {
	a := []int{4, 2, 7, 9, 1, 8, 5, 6, 3}
	QuickSort(a)
	fmt.Println(a)
}

func TestSum(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := SumV2(a)
	fmt.Println(b)
}
func TestFact(t *testing.T) {
	n := 35
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		s := time.Now()
		defer fmt.Println("递归耗时:", time.Since(s))
		fmt.Println(Fact(n))
		wg.Done()
	}()
	go func() {
		s := time.Now()
		defer fmt.Println("尾递归耗时:", time.Since(s))
		fmt.Println(FactTail(n, 1))
		wg.Done()
	}()
	wg.Wait()
}
func TestFib(t *testing.T) {
	n := 7
	fmt.Println(Fib(n))
	fmt.Println(FibByFor(n))
}

func BenchmarkStep(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		Step(10)
	}
}
func BenchmarkStepMap(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		StepMap(10)
	}
}

func BenchmarkStepF(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		StepF(10)
	}
}