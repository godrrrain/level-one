package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов с использованием конкурентных вычислений.
*/

func justChannels(arr []int) int {
	messages := make(chan int, len(arr))

	for _, v := range arr {
		go func(val int) {
			messages <- val * val
		}(v)
	}

	sum := 0
	for range arr {
		sum += <-messages
	}
	return sum
}

func waitGroups(arr []int) int {
	var wg sync.WaitGroup

	sum := 0
	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			sum += val * val
			wg.Done()
		}(val)
	}
	wg.Wait()
	return sum
}

func atomicS(arr []int) int {
	var (
		wg     sync.WaitGroup
		result int64
	)

	for _, v := range arr {
		wg.Add(1)
		go func(num int64) {
			defer wg.Done()
			// потокобезопасное выполнение
			atomic.AddInt64(&result, num*num)
		}(int64(v))
	}

	wg.Wait()

	return int(result)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Printf("Ch: %d\n", justChannels(arr))
	fmt.Printf("Wg: %d\n", waitGroups(arr))
	fmt.Printf("At: %d\n", atomicS(arr))
}
