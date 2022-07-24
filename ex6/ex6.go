package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		Способ №1 WaitGroup
	*/
	fmt.Println("WaitGroup start")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go waitGroupeStop(&wg)
	wg.Wait()
	fmt.Println("WaitGroup stop")
	/*
		Способ №2 Context
	*/
	fmt.Println("Context start")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg1 := sync.WaitGroup{}
	wg1.Add(1)
	go contextStop(ctx)
	time.Sleep(time.Microsecond * 30)
	cancel()
	fmt.Println("Context stop")

}

func waitGroupeStop(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	defer wg.Done()
}

func contextStop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Doing something")
		}
	}
}
