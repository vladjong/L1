package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	Способ №3 Context.WithDeadline
*/
func main() {
	fmt.Println("Context.WithDeadline start")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer cancel()
	go contextStop(ctx, &wg)
	wg.Wait()
	fmt.Println("Context.WithDeadline stop")

}

func contextStop(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Doing something")
		}
	}
}
