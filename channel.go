package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var Channel = make(chan string)

func InputData() {
	fmt.Println("Ready to put datas")
	for i := 0; i < 10; i++ {
		Channel <- "数据" + string(i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func OutputData() {
	fmt.Println("Prepare to get datas")
	for v := range Channel {
		fmt.Println(v)
		//time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func main() {
	go InputData()
	go OutputData()
	wg.Add(2)
	wg.Wait()
}
