package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(arr []int, target int) int {
	//границы поиска
	low, high := 0, len(arr)-1

	for low <= high {

		mid := low + (high-low)/2 //средний индекс

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			// если искомый элемент больше среднего, переходим к правой части слайса
			low = mid + 1
		} else {
			// если искомый элемент меньше среднего, переходим к левой части слайса
			high = mid - 1
		}
	}
	// если элемент не найден, возвращаем -1
	return -1
}

func main() {
	arr := []int{-10, -8, -5, 0, 3, 5, 8, 10, 16}
	target := 8
	resIndex := binarySearch(arr, target)
	if resIndex == -1 {
		fmt.Printf("Элемент %d не найден", target)
	} else {
		fmt.Printf("Элемент %d имеет индекс %d", target, resIndex)
	}
}
