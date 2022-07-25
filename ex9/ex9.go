package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array:", arr)
	/*
		Иницилизируем два канала
	*/
	writeNumber := make(chan int)
	readNumber := make(chan int)
	/*
		Иницилизируем waitGroup для синхронной работы горутин
	*/
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go readArray(writeNumber, arr)
	go multiply(writeNumber, readNumber)
	go read(readNumber, wg)
	wg.Wait()
}

/*
	Функция чтения чисел из массива и записи в writeNumber канал
*/
func readArray(writeNumber chan int, arr []int) {
	for _, val := range arr {
		writeNumber <- val
	}
	/*
		Закрываем канал, после того как прочитали все данные из массива
	*/
	close(writeNumber)
}

/*
	Функция чтения чисел из канала и запись во второй канал
*/
func multiply(writeNumber chan int, readNumber chan int) {
	for item := range writeNumber {
		readNumber <- item * 2
	}
	/*
		Закрываем канал, если канал writeNumber закрыт
	*/
	close(readNumber)
}

/*
	Функция вывод чисел на экран
*/
func read(readNumber chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range readNumber {
		fmt.Println("number:", item)
	}
}
