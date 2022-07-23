package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	/*Иницилизируем генератор рандомных чисел*/
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	/*Считываем количество воркеры*/
	fmt.Println("Enter number of workers: ")
	cntWorker := 0
	fmt.Scan(&cntWorker)
	/*Создаем новый канал*/
	ch := make(chan int)
	/*Создаем новую группу, в которую добавляем кол-во всех воркеров*/
	wg := sync.WaitGroup{}
	wg.Add(cntWorker)
	fmt.Println("Start program")
	/*запускаем воркеров*/
	initWorker(cntWorker, ch, &wg)
	/*запись данных в канал + проверка на сигнал*/
	sendValueInChanel(ch, r, &wg)
}

func initWorker(cntWorker int, ch chan int, wg *sync.WaitGroup) {
	for w := 1; w <= cntWorker; w++ {
		fmt.Println("Worker", w, "go work!")
		/*функция чтения данных из канала и передача в главный поток*/
		go worker(w, ch, wg)
		time.Sleep(time.Millisecond * 100)
	}
}

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	/*сигнализирует что воркер завершил выполнение*/
	defer wg.Done()
	/*Чтение записи из канала в главном потоке*/
	for value := range ch {
		fmt.Println(value)
	}
	/*Оповещение о закрытие канала*/
	fmt.Printf("Worker ID: %d, closed\n", id)
	/*Оповещение о закрытие канала*/
}

func sendValueInChanel(ch chan int, r *rand.Rand, wg *sync.WaitGroup) {
	sigs := make(chan os.Signal, 1)
	/*подписываемся на сигнал CTR+C*/
	signal.Notify(sigs, syscall.SIGINT)
	for {
		select {
		/*кейс CTR+C*/
		case signal := <-sigs:
			fmt.Println("\nSignal", signal, "received")
			/*закрываем канал*/
			close(ch)
			/*дожидаемся окончание работы воркеров*/
			wg.Wait()
			return
		default:
			/*запись в канал*/
			ch <- r.Intn(10)
			time.Sleep(time.Millisecond * 100)
		}
	}
}
