package main

import (
	"fmt"
	"sync"
)

func pow(element int) int {
	return element * element
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	sum := 0
	for _, item := range arr {
		wg.Add(1)
		go func(item int) {
			result := pow(item)
			sum += result
			wg.Done()
		}(item)
		wg.Wait()
	}
	fmt.Println(sum)
}
