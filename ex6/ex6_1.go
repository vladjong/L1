package main

import (
	"fmt"
	"sync"
)

/*
	Способ №1 WaitGroup
*/
func main() {
	fmt.Println("WaitGroup start")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go waitGroupeStop(&wg)
	wg.Wait()
	fmt.Println("WaitGroup stop")
}

func waitGroupeStop(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	defer wg.Done()
}
