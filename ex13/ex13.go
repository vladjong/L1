package main

import "fmt"

func main() {
	a := 27
	b := 7
	fmt.Println("Before a:", a, "b:", b)
	first(a, b)
	second(a, b)
	third(a, b)
}

/*
	Средства языка
*/
func first(a int, b int) {
	a, b = b, a
	fmt.Println("First a:", a, "b:", b)
}

/*
	Матемачиские фокусы
*/
func second(a int, b int) {
	a += b
	b = a - b
	a -= b
	fmt.Println("Second a:", a, "b:", b)
}

/*
	XOR
*/
func third(a int, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("Second a:", a, "b:", b)
}
