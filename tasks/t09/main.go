package main

import (
	"fmt"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

func insertData(numArr []int, inputCh chan int) {
	for _, num := range numArr {
		inputCh <- num
	}
	close(inputCh)
}

func changeData(inputCh chan int, outputCh chan int) {
	for num := range inputCh {
		res := num * 2
		outputCh <- res
	}
	close(outputCh)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	inputCh := make(chan int)
	outputCh := make(chan int)

	go insertData(numbers, inputCh)
	go changeData(inputCh, outputCh)

	for res := range outputCh {
		fmt.Println(res)
	}
}
