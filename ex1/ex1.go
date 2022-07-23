package main

import "fmt"

type Human struct {
	age    int
	height float64
}

func (h *Human) say() {
	fmt.Println("Hello my friend's!")
}

func (h *Human) do(action string) {
	fmt.Println("I", action)
}

func (h *Human) information() string {
	return fmt.Sprintf("Some info about me: age = %d, height = %.2f", h.age, h.height)
}

type Action struct {
	Human
	name string
}

func main() {
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
