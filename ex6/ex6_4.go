package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Способ №3 Context.WithTimeout
*/
func main() {
	fmt.Println("Context.WithTimeout start")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go contextStop(ctx)
	time.Sleep(time.Second)
	fmt.Println("Context.WithTimeout stop")
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
