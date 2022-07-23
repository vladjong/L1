package main

import "fmt"

func main() {
	arr := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Println("Source array:", arr)
	fmt.Println("Enter position delete: ")
	pos := 0
	fmt.Scan(&pos)
	if pos >= len(arr) {
		fmt.Println("Incorect index")
		return
	}
	arr1, arr2 := arr, arr
	arr1 = deleteElementChangeOrder(pos, arr1)
	fmt.Println("Changer array:", arr1)
	arr2 = deleteElementySaveOrder(pos, arr)
	fmt.Println("Changer array:", arr2)
}

/*
	Способ с изменением массива
*/
func deleteElementChangeOrder(pos int, arr []string) []string {
	/*
		Меняем последний символ местами с pos
	*/
	arr[pos] = arr[len(arr)-1]
	/*
		Делаем слайс по последнему элементу
	*/
	arr = arr[:len(arr)-1]
	return arr
}

/*
	Способ с сохранением порядка в массиве
*/
func deleteElementySaveOrder(pos int, arr []string) []string {
	/*
		Делаем два слайса до pos и после, далее складываем их
	*/
	arr = append(arr[:pos], arr[pos+1:]...)
	return arr
}
