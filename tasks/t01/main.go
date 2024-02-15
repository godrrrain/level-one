package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

type Human struct {
	Name string
	Age  uint8
	City string
}

func (h *Human) setName(name string) {
	h.Name = name
}

func (h *Human) setAge(age uint8) {
	h.Age = age
}

func (h *Human) setCity(city string) {
	h.City = city
}

func (h *Human) ShowInfoHuman() {
	fmt.Printf("[Human] Name: %s, Age: %d, City: %s.\n", h.Name, h.Age, h.City)
}

type Action struct {
	Human
}

func (a *Action) ShowInfoAction() {
	fmt.Printf("[Action] Name: %s, Age: %d, City: %s.\n", a.Name, a.Age, a.City)
}

func NewAction(name string, age uint8, city string) *Action {
	a := Action{}

	// Вызываем родительские методы Human
	a.setName(name)
	a.setAge(age)
	a.setCity(city)

	return &a
}

func main() {
	action := NewAction("Vova", 22, "Moscow")

	// Вызываем метод структуры Action
	action.ShowInfoAction()

	// Вызываемм метод структуры Human
	action.ShowInfoHuman()

}
