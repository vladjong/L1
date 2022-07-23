package main

import "fmt"

func main() {
	fmt.Println("Enter number:")
	var number int64
	fmt.Scan(&number)
	fmt.Println("Choose position bit:")
	var bitPosition int64
	fmt.Scan(&bitPosition)
	fmt.Println("Select the bit value 1 or 0:")
	var bitValue int64
	fmt.Scan(&bitValue)
	if bitValue == 1 {
		/*
			Операция or true
		*/
		number = number | (1 << bitPosition)
	} else {
		/*
			Операция xor true
		*/
		number = number ^ (1 << bitPosition)
	}
	fmt.Println("Result value:", number)
}
