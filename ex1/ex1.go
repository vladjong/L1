package main

import "fmt"

/*
	Создаем структуру Human
*/
type Human struct {
	age    int
	height float64
}

/*
	Методы структуры Human
*/
func (h *Human) say() {
	fmt.Println("Hello my friend's!")
}

func (h *Human) do(action string) {
	fmt.Println("I", action)
}

func (h *Human) information() string {
	return fmt.Sprintf("Some info about me: age = %d, height = %.2f", h.age, h.height)
}

/*
	Создаем структуру Action.
	Унаследуемся от структуры Human,
	для этого в структуру Action вложим родительскую структуру Human
*/
type Action struct {
	Human
	name string
}

func main() {
	/*
		Чтобы создать экземпляр структуры Action,
		нужно иницилизировать все поля ролительской структуры,
		а также свои поля
	*/
	action := Action{
		Human: Human{
			age:    23,
			height: 194.3,
		},
		name: "learn Go",
	}
	action.say()
	info := action.information()
	fmt.Println(info)
	action.do(action.name)
}
