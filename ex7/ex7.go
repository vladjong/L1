package main

import (
	"fmt"
	"sync"
)

/*
	Создадим структуру Map:
	1. Указатель на мьютекс, чтобы использовать методы мьютекса в самой структуре
	2. Создадим map: key - int, value - string
*/
type Map struct {
	mapMutex *sync.Mutex
	myMap    map[int]string
}

/*
	Метод записи в Map. Перед записью останавливаем Lock() все горутины,
	после вставки возобновляем работу методом Unlock()
*/
func (m *Map) Write(key int, str string, wg *sync.WaitGroup) {
	defer wg.Done()
	m.mapMutex.Lock()
	m.myMap[key] = str
	m.mapMutex.Unlock()
}

func main() {
	obj := Map{
		myMap:    make(map[int]string, 0),
		mapMutex: new(sync.Mutex),
	}
	/*
		Создадим WaitGroup и записываем туда 10 заданий методом Add
	*/
	wg := sync.WaitGroup{}
	wg.Add(10)
	/*
		Запускаем в работу наши горутины
	*/
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("Gorutin %d", i)
		go obj.Write(i, str, &wg)
	}
	/*
		Метод Wait ждет пока завершаться все задания в группе
	*/
	wg.Wait()
	for key, value := range obj.myMap {
		fmt.Println("Key:", key, "Value: ", value)
	}
}
