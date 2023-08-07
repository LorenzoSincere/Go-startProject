package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//for是主线程
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	//子协程执行完之前主线程不退出
	time.Sleep(time.Second)

	CalSquare()

	//优化
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}

func hello(j int) {
	println("hello goroutine : " + fmt.Sprint(j))
}

func CalSquare() {
	src := make(chan int)
	//解决生产消费速度匹配问题
	dest := make(chan int, 3)
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()

	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()
	for i := range dest {
		println(i)
	}
}
