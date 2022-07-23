package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Enter time duration of work:")
	choseTime := 0
	fmt.Scan(&choseTime)
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go read(ch, &wg)
	write(ch, choseTime, &wg)
}

func read(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		fmt.Println(item)
	}
}

func write(ch chan string, choseTime int, wg *sync.WaitGroup) {
	timeout := time.After(time.Duration(choseTime) * time.Second)
	for {
		select {
		case <-timeout:
			fmt.Println("My time is up\nGoodbye!")
			close(ch)
			wg.Wait()
			return
		default:
			ch <- "Hello"
			time.Sleep(time.Millisecond * 100)
		}
	}
}
