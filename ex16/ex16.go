package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := generateSlice()
	fmt.Println("Not sort slice: ", arr)
	arr = quicksort(arr)
	fmt.Println("Sort slice: ", arr)
}

func quicksort(arr []int) []int {
	/*Базовый случай*/
	if len(arr) < 2 {
		return arr
	}
	/*Иницилизируем два указателя*/
	left, right := 0, len(arr)-1
	/*Выбираем рандомный разделитель*/
	pivot := rand.Int() % len(arr)
	/*Меняем правый элемент с выбранным*/
	arr[pivot], arr[right] = arr[right], arr[pivot]
	/*Иттерируемся по массиву, методом двух указателей*/
	for i, _ := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	/*Меняем правый элемент с левым*/
	arr[left], arr[right] = arr[right], arr[left]
	/*Запускаем алгоритм для левой половины*/
	quicksort(arr[:left])
	/*Запускаем алгоритм для правой половины*/
	quicksort(arr[left+1:])
	return arr
}

func generateSlice() []int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	size := r.Intn(50)
	var arr []int = make([]int, int(size))
	for i := 0; i < size; i++ {
		arr[i] = r.Intn(100)
	}
	return arr
}
