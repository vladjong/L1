package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	Структура Counter
	1. value - значение
	2. sync.Mutex - мьютекс
*/
type Counter struct {
	value int
	sync.Mutex
}

func main() {
	counter := Counter{}
	fmt.Println("Enter time duration: ")
	cntTime := 0
	fmt.Scan(&cntTime)
	/*
		Создаем контекст, который истекает через введенное время пользователем
	*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cntTime)*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	/*
		Устанавливаем число логичечских ядер на процессоре
	*/
	numCPU := runtime.NumCPU()
	/*
		Добавляем в группу количество заданий равное ядрам процессора
	*/
	wg.Add(numCPU)
	/*
		Создаем горутину
	*/
	for i := 0; i < numCPU; i++ {
		go countNumber(ctx, &wg, &counter)
	}
	/*
		Ждем когда выполнятся все наши горутины
	*/
	wg.Wait()
	fmt.Println("Counter:", counter.value)
}

/*
	Функция увелечения счетчика на 1
*/
func countNumber(ctx context.Context, wg *sync.WaitGroup, counter *Counter) {
	/*
		Функция увелечения счетчика на 1
	*/
	defer wg.Done()
	for {
		select {
		/*
			Кейс, когда контекст завершается
		*/
		case <-ctx.Done():
			return
		/*
			Инкрементируем наш счетчик.
			Используем мьютекс, для блокировки горутин.
		*/
		default:
			counter.Mutex.Lock()
			counter.value++
			counter.Mutex.Unlock()
		}
	}
}
