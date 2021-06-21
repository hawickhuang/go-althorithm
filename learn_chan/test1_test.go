package learn_chan

import (
	"fmt"
	"testing"
	"time"
)

func Test_goroutine(t *testing.T) {
	ch := make(chan int, 3)
	go goroutineA(ch)
	go goroutineB(ch)
	ch <- 3
	time.Sleep(time.Second * 1)
}

func Test_pc(t *testing.T) {
	c := make(chan int)
	isClosed := make(chan bool)
	pClosed := make(chan bool)
	defer close(isClosed)
	defer close(pClosed)
	defer close(c)

	go producer(c, pClosed, isClosed)
	go consumer(c, pClosed)

	<-isClosed
	fmt.Printf("done")
}

func Test_noInit(t *testing.T) {
	var c chan int
	//go func() {
	//	c <- 1
	//}()
	fmt.Println(<-c)
}

func Test_channel1(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	t.Logf("x = %d", <-c)
}

func Test_channel2(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		c1 <- 1
	}()
	go func() {
		c2 <- "string"
	}()

	for {
		select {
		case data, ok := <-c1:
			t.Logf("receive from c1, data: %d, ok: %t", data, ok)
		case data, ok := <-c2:
			t.Logf("receive from c2, data: %s, ok: %t", data, ok)
		case <-time.After(time.Second*3):
			t.Logf("timeout")
			return
		}
	}
}

func Test_channel3(t *testing.T) {
	c := make(chan int)
	go func() {
		for i:=0; i<10;i++{
			c <- i
		}
		//close(c)
	}()

	go func() {
		for i := range c {
			t.Logf("data: %d", i)
		}
	}()
	time.Sleep(time.Second)
}
