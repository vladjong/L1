package main

import (
	"fmt"
)

var justString string

func someFunc() {
	/*
		Функция возвращает []rune/string из 1024 элементов
	*/
	v := createHugeString(1 << 10)
	fmt.Println(v)
	/*
		Строка в Go реализована на массиве из байт и длины строки.
		Если мы создаем слайс из строки, то она указывает на исходную строку.
		Решение: вместо исходной строки будем сохранять []byte и приводить к типу,
		чтобы создать копию нашей строки.
	*/
	justString = string([]byte(v[:100]))
	fmt.Println(justString)
}

func main() {
	someFunc()
}

func createHugeString(count int) string {
	var s string
	for i := 0; i < count; i++ {
		s += "g"
	}
	return s
}
