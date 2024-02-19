package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере
*/

// целевой интерфейс
type Target interface {
	Request() string
}

// адаптируемый класс
type Adaptee struct {
}

func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// адаптер
type Adapter struct {
	adaptee *Adaptee
}

func (adapter *Adapter) Request() string {
	return adapter.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}

	// клиентский код работает через интерфейс Target
	fmt.Println(adapter.Request())
}
