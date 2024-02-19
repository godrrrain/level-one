package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
*/

// структура счетчика
type safeCounter struct {
	mutex sync.Mutex
	value int
}

func (sc *safeCounter) Increment() {
	sc.mutex.Lock()         // остановка других горутин, пытающихся вызвать метод Lock()
	defer sc.mutex.Unlock() // продолжение работы других горутин после завершения текущей

	sc.value++
}

func newSafeCounter() *safeCounter {
	return &safeCounter{
		sync.Mutex{},
		0,
	}
}

func (sc *safeCounter) GetValue() int {
	return sc.value
}

func main() {
	counter := newSafeCounter()

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	// создание воркеров
	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func(counter *safeCounter, ctx context.Context, wg *sync.WaitGroup) {
			for {
				counter.Increment()
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
				}

				time.Sleep(time.Second)
			}

		}(counter, ctx, &wg)
	}

	time.Sleep(time.Second * 5)

	cancel()
	wg.Wait()

	result := counter.GetValue()
	fmt.Println("Counter result:", result)
}
