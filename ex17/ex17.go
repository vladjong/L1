package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("Enter pattern: ")
	pattern := 0
	fmt.Scan(&pattern)
	arr := generateSlice()
	sort.Ints(arr)
	fmt.Println("Given array: ", arr)
	position := binarySearch(arr, pattern)
	if position == -1 {
		fmt.Println("The pattern will not find")
	} else {
		fmt.Println("The pattern position: ", position)
	}
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

/*
	Алгоритм бинарного поиска
*/
func binarySearch(arr []int, pattern int) int {
	max, min := len(arr), 0
	for i := 0; i < len(arr); i++ {
		index := (max - min) / 2
		if arr[index] == pattern {
			return index
		} else if arr[index] < pattern {
			min = index
		} else {
			max = index
		}
	}
	return -1
}
