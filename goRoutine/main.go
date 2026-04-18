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
		time.Sleep(time.Millisecond * 100)
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

// 互斥锁
var mutex sync.Mutex
var v int

func ComputingAddUnit(position int) {
	wg.Add(1)
	//mutex.Lock()
	v += 1
	fmt.Printf("======位次：%d 的增V:%d\n", position, v)
	//time.Sleep(time.Millisecond * 200)
	//mutex.Unlock()
	wg.Done()
}
func ComputingDesUnit(position int) {
	wg.Add(1)
	v -= 1
	time.Sleep(time.Millisecond * 2000)
	fmt.Printf("位次：%d 的减V:%d\n", position, v)

	wg.Done()
}

// 读写锁
var mutexRW sync.RWMutex

func main() {
	//wg.Add(2)
	//go thread()
	//wg.Add(1)
	//go thread2()
	//wg.Wait()
	//wg.Add(2)
	//go channelSend()
	//go channelRecv()
	//wg.Wait()
	for i := 0; i < 50; i++ {
		go ComputingAddUnit(i)
		go ComputingDesUnit(i)
	}
	wg.Wait()

}
