package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами
языка
*/

func main() {
	arr := []int{1, -2, 10, 8, 5, 6, -5, -10}

	fmt.Println("Изначальный массив: ", arr)

	quicksort(arr, 0, len(arr)-1)

	fmt.Println("Отсортированный массив: ", arr)
}

func quicksort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quicksort(arr, low, pivotIndex-1)
		quicksort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
