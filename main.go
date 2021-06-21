package main

import (
	"fmt"
)

func main()  {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	x := <-c

	fmt.Println("x=", x)
}
