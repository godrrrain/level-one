package main

import (
	"fmt"
	"slices"
)

/*
Удалить i-ый элемент из слайса.
*/

func main() {
	sampleSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Изначальные слайс:", sampleSlice)

	// получение индекса для удаления из слайса
	var index int
	fmt.Print("Введите индекс элемента для удаления: ")
	_, err := fmt.Scan(&index)
	if err != nil {
		fmt.Println("Введено некорректное число")
		return
	}
	fmt.Println("Используя slices.Delete():")
	sampleSlice = slices.Delete(sampleSlice, index, index+1)
	fmt.Println(sampleSlice)

	fmt.Println("Используя реслайсинг:")
	sampleSlice = []int{1, 2, 3, 4, 5}
	sampleSlice = append(sampleSlice[:index], sampleSlice[index+1:]...) // переопределение слайса на составленный из двух частей старого - до индекса удаления и после него
	fmt.Println(sampleSlice)

	fmt.Println("Неупорядоченный реслайсинг:")
	sampleSlice = []int{1, 2, 3, 4, 5}
	sampleSlice[index] = sampleSlice[len(sampleSlice)-1] // приравнивание элемента с индексом удаления к последнему
	sampleSlice = sampleSlice[:len(sampleSlice)-1]       // переопределение слайса на старый без последнего элемента

	fmt.Println(sampleSlice)
}
