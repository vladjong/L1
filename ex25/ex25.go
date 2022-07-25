package main

import (
	"fmt"
	"time"
)

/*
	Реализация функции sleep
*/
func sleep(value int) {
	select {
	/*
		Когда время истекает, то канал возвращает значение текущекого времени и завершаем работу.
	*/
	case <-time.After(time.Duration(value) * time.Second):
		return
	}
}

func main() {
	value := 0
	fmt.Println("Enter time duration sleep: ")
	fmt.Scan(&value)
	fmt.Println("Sleep", time.Now().Second())
	sleep(value)
	fmt.Println("Wake up", time.Now().Second())
}
