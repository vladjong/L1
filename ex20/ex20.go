package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	strArr := strings.Split(str, " ")
	strP := palindrom(strArr)
	fmt.Println(str, "—", strP)
}

/*Метод двух указателей*/
func palindrom(strArr []string) string {
	left, right := 0, len(strArr)-1
	for left < right {
		strArr[left], strArr[right] = strArr[right], strArr[left]
		left++
		right--
	}
	return strings.Join(strArr, " ")
}
