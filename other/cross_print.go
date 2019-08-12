package other

import (
	"fmt"
	"time"
)

func WaitChan() {
	exit := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		fmt.Println("print go")
		close(exit)
	}()
	<-exit
	fmt.Println("exit")
}

func PrintString() {
	str := "abcdefg"
	str1 := "gfedcba"
	fmt.Println(str[:1])
	fmt.Println(str1[:1])
	c1 := make(chan bool, 1)
	c2 := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-c1:
				fmt.Println("a")
				c2 <- true
			}
		}
	}()
	go func() {
		for {
			select {
			case <-c2:
				fmt.Println("b")
				c1 <- true
			}
		}
	}()
	time.Sleep(time.Second)

}
