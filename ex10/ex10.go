package main

import "fmt"

func main() {
	arr := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	setOrNumber := groupNumber(arr)
	fmt.Println(setOrNumber)
}

/*метод объединения в группы*/
func groupNumber(arr [8]float64) map[int][]float64 {
	setOrNumber := make(map[int][]float64)
	for _, value := range arr {
		/*преобразуем элемент массива в int и умнажаем на 10, чтобы получить шаг 10*/
		var key int = int(value) / 10 * 10
		setOrNumber[key] = append(setOrNumber[key], value)
	}
	return setOrNumber
}
