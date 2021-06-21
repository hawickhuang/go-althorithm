package learn_chan

import (
	"fmt"
	"math/rand"
)

func goroutineA(a <-chan int) {
	val := <-a
	fmt.Println("goroutineA received the data: ", val)
}

func goroutineB(b chan int) {
	val := <-b
	fmt.Println("goroutineB received the data: ", val)
}

func producer(c chan<- int, pc <-chan bool, isClosed chan<- bool) {
	for {
		select {
		case <-pc:
			fmt.Println("producer closed")
			isClosed <- true
			return
		default:
			c <- rand.Int()
		}
	}
}

func consumer(c <-chan int, pClosed chan<- bool) {
	for i := 0; i < 5; i++ {
		val := <-c
		fmt.Println("recevie: ", val)
	}
	pClosed <-true
}

