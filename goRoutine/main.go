package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func thread() {
	for i := 0; i < 10; i++ {
		fmt.Printf("thread1 %d\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}
func thread2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("thread2 %d", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func channel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

var ch = make(chan int, 10)

func channelSend() {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 300)
	}
	close(ch)
	wg.Done()
}
func channelRecv() {
	for v := range ch {
		fmt.Printf("收到管道数据：%d\n", v)
	}
	wg.Done()
}

func main() {
	//wg.Add(2)
	//go thread()
	//wg.Add(1)
	//go thread2()
	//wg.Wait()
	wg.Add(2)
	go channelSend()
	go channelRecv()
	wg.Wait()

}
