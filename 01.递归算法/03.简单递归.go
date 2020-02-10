package Recurive

import (
	"fmt"
	"runtime/debug"
)

func Print(i, j int) {
	if i < j {
		mid := (i + j) / 2
		fmt.Printf("a: i:%d, j:%d, mid:%d\n", i, j, mid)
		print(i, mid)
		print(mid+1, j)
		fmt.Printf("b: i:%d, j:%d, mid:%d\n", i, j, mid)
		debug.PrintStack()

	}
}

