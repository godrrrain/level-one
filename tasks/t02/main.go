package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func justChannels(arr []int) {
	messages := make(chan int, len(arr))

	for _, v := range arr {
		go func(val int) {
			messages <- val
		}(v)
	}

	for range arr {
		val := <-messages
		fmt.Printf("Ch: %d\n", val*val)
	}
}

func waitGroups(arr []int) {
	var wg sync.WaitGroup
	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			fmt.Printf("Wg: %d\n", val*val)
			wg.Done()
		}(val)
	}

	wg.Wait()
}

func main() {

	arr := []int{2, 4, 6, 8, 10}
	justChannels(arr)
	waitGroups(arr)
}
