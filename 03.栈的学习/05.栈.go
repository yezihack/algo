package main

import (
	"code.qschou.com/golang/errors"
	"fmt"
	"log"
)

func main() {
	stack := make([]int, 0, 5)
	push := func(x int) error {
		if len(stack) >= cap(stack) {
			return errors.New("stack is full")
		}
		stack = append(stack, x)
		return nil
	}
	pop := func() (x int, err error) {
		if len(stack) == 0 {
			err = errors.New("stack is empty")
			return
		}
		x = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return
	}
	for i := 1; i <= 5; i ++ {
		if err := push(i); err != nil {
			log.Fatal(err)
			break
		}
	}
	for i := 0;i < 5; i++ {
		x, err := pop()
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("x:%d\n", x)
	}
}
