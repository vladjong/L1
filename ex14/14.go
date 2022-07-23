package main

import (
	"fmt"
	"reflect"
)

func main() {
	var first interface{} = 21
	var second interface{} = "Go"
	var third interface{} = true
	var fourth interface{} = make(chan int)
	determinateDataSwitch(first)
	determinateDataSwitch(second)
	determinateDataSwitch(third)
	determinateDataSwitch(fourth)
	determinateDataReflect(first)
	determinateDataReflect(second)
	determinateDataReflect(third)
	determinateDataReflect(fourth)
}

func determinateDataSwitch(x interface{}) {
	switch data := x.(type) {
	case int:
		fmt.Println("int", data)
	case string:
		fmt.Println("string", data)
	case bool:
		fmt.Println("bool", data)
	case chan int:
		fmt.Println("chan int", data)
	default:
		fmt.Println("Unknown data type")
	}
}

/*
	Пакет reflect содержит метод TypeOf для определения типа переменной
*/
func determinateDataReflect(x interface{}) {
	fmt.Println(reflect.TypeOf(x), x)
}
