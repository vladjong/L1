package main

import "fmt"

func main() {
	fmt.Println("Enter word:")
	var str string
	fmt.Scanln(&str)
	strP := palindrom(str)
	fmt.Println(str, "—", strP)
}

/*Метод двух указателей*/
func palindrom(str string) string {
	/*Преобразуем строку в массив рун, чтобы работала с unicode*/
	strByte := []rune(str)
	left, right := 0, len(strByte)-1
	for left < right {
		strByte[left], strByte[right] = strByte[right], strByte[left]
		left++
		right--
	}
	return string(strByte)
}
