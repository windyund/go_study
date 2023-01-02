package main

import (
	"fmt"
	"time"
)

func task(a int, b int) int {
	return a + b
}

func main() {
	c := make(chan int, 3)
	go func() {
		fmt.Println("sub go routine run")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("goroutine put value=", i)
		}
		//close(c)
	}()

	time.Sleep(1 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
	res := <-c
	fmt.Println("main recevie res:", res)

}
