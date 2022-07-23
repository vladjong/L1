package main

import "fmt"

func main() {
	first := [9]int{2, 1, 3, 43, 23, 6, 4, 8, 9}
	second := [9]int{1, 22, 3, 6, 5, 8, 7, 4, 9}
	intersection := intersectionSet(first, second)
	fmt.Println(intersection)
}

/*
	Создаем мапу, если ключ уже встречался, то прибалвяем 1. Те у кого 2 являются пересечением
*/
func intersectionSet(first [9]int, second [9]int) []int {
	valOrInter := make(map[int]int)
	for _, val := range first {
		valOrInter[val] += 1
	}
	for _, val := range second {
		valOrInter[val] += 1
	}
	intersection := []int{}
	for key, val := range valOrInter {
		if val == 2 {
			intersection = append(intersection, key)
		}
	}
	return intersection
}
