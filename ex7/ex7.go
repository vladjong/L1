package main

import (
	"fmt"
	"sync"
)

type Map struct {
	mapMutex *sync.Mutex
	myMap    map[int]string
}

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
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("Gorutin %d", i)
		go obj.Write(i, str, &wg)
	}
	wg.Wait()
	for key, value := range obj.myMap {
		fmt.Println("Key:", key, "Value: ", value)
	}
}
