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
	/*
		Создадим WaitGroup, и запишем туда в цикле кол-во заданий методом Add.
	*/
	var wg sync.WaitGroup
	sum := 0
	for _, item := range arr {
		wg.Add(1)
		/*
			Реализуем вложенную функцию, по завершению которой будет вызывать метод Done у WaitGroup
		*/
		go func(item int) {
			result := pow(item)
			sum += result
			wg.Done()
		}(item)
		/*
			Метод Wait ждет пока завершаться все задания в группе
		*/
		wg.Wait()
	}
	fmt.Println(sum)
}
