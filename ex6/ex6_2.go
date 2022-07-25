package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Способ №2 Context.WithCancel
*/
func main() {
	fmt.Println("Context.WithCancel start")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go contextStop(ctx)
	time.Sleep(time.Microsecond * 50)
	cancel()
	fmt.Println("Context.WithCancel stop")
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
