package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Enter string: ")
	var str string
	fmt.Scan(&str)
	str = strings.ToLower(str)
	sortMethod(str)
	mapMethod(str)
}

/*
	Сортируем и смотрим на дубликаты
*/
func sortMethod(str string) {
	r := []rune(str)
	sort.Slice(r, func(i int, j int) bool { return r[i] < r[j] })
	str = string(r)
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			fmt.Println("First: true")
			return
		}
	}
	fmt.Println("First: false")
}

/*
	Map хранит только уникальные элементы
*/
func mapMethod(str string) {
	runeMap := make(map[rune]bool)
	for _, item := range str {
		_, found := runeMap[item]
		if found == true {
			fmt.Println("Second: true")
			return
		}
		runeMap[item] = true
	}
	fmt.Println("Second: false")
}
