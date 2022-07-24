package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Способ №1 с использованием time.After.
	After ожидает истечение продолжительности времени,
	а затем отправляет текущее время в канал
*/
func main() {
	fmt.Println("Enter time duration of work:")
	choseTime := 0
	fmt.Scan(&choseTime)
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go read(ch, &wg)
	go write(ch, choseTime)
	wg.Wait()
}

/*
	Функция чтение записи из канала,
	перебираем все значения полученные из канала
*/
func read(ch chan string, wg *sync.WaitGroup) {
	for item := range ch {
		fmt.Println(item)
	}
	defer wg.Done()
}

/*
	Функция записи в канал, если время истекает,
	то мы переходмс в кейс timeout, закрываем канал.
*/
func write(ch chan string, choseTime int) {
	timeout := time.After(time.Duration(choseTime) * time.Second)
	for {
		select {
		case <-timeout:
			fmt.Println("My time is up\nGoodbye!")
			close(ch)
			return
		default:
			ch <- "Hello"
			time.Sleep(time.Millisecond * 100)
		}
	}
}

/*
	Способ №2 с использованием context.WithTimeout.
	After ожидает истечение продолжительности времени,
	а затем отправляет текущее время в канал
*/

/*
func main() {
	fmt.Println("Enter time duration of work:")
	choseTime := 0
	fmt.Scan(&choseTime)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	wg := new(sync.WaitGroup)
	ch := make(chan string)
	wg.Add(1)
	go write(ctx, ch)
	go reader(ch, wg)
	wg.Wait()
}

func write(ctx context.Context, ch chan<- string) {
	for {
		select {
		default:
			ch <- "Hello"
			time.Sleep(time.Millisecond * 100)
		case <-ctx.Done():
			fmt.Println("My time is up\nGoodbye!")
			close(ch)
			return
		}
	}
}
*/
