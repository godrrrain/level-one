package main

import (
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	var num int64
	var i uint

	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Введите корректное число")
		return
	}
	fmt.Print("Введите позицию бита: ")
	_, err = fmt.Scan(&i)
	if err != nil {
		fmt.Println("Введите корректное число")
		return
	}

	// Установка i-го бита в 1
	num = num | (1 << i)
	fmt.Println("Число после установки бита в 1:", num)

	// Установка i-го бита в 0
	num = num & ^(1 << i)
	fmt.Println("Число после установки бита в 0:", num)
}
