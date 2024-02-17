package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func main() {
	var wg sync.WaitGroup
	m := make(map[int]int)
	mutex := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(mutex *sync.Mutex) {
			defer wg.Done()

			index := rand.Intn(100)
			value := rand.Intn(100)
			// Ограничиваем доступ остальных горутин на запись
			mutex.Lock()
			m[index] = value
			mutex.Unlock()

		}(&mutex)
	}

	wg.Wait()

	// Выводим значения из карты
	for k, v := range m {
		fmt.Printf("Ключ: %d, Значение: %d\n", k, v)
	}
}
