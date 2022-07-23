package main

import (
	"fmt"
	"math/big"
)

func main() {
	/*
		a = 21e31 и b ~ 12e27
	*/
	a := new(big.Int)
	a.SetString("210000000000000000000000000000000", 10)
	b := new(big.Int)
	b.SetString("12000000230000420000002402423", 10) // 56 * 10^32
	/*
		Сложение
	*/
	c := new(big.Int)
	c.Add(a, b)
	fmt.Println("a + b =", c)
	/*
		Вычитание
	*/
	c.Sub(a, b)
	fmt.Println("a - b =", c)
	/*
		Умножение
	*/
	c.Mul(a, b)
	fmt.Println("a * b =", c)
	/*
		Деление
	*/
	c.Div(a, b)
	fmt.Println("a / b =", c)
}
