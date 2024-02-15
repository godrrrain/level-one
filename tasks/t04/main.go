package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	channel := make(chan int, 10)

	var n int
	fmt.Print("Enter the number of workers: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Enter a valid number")
		return
	}

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background()) // создание контекста и объявление функции завершения

	// создание воркеров
	for i := 0; i < n; i++ {

		wg.Add(1)
		// запуск горутины воркера
		go func(wn int, ctx context.Context, wg *sync.WaitGroup) {
			for {
				// неблокирующее получение данных из канала
				select {
				case data := <-channel:
					fmt.Println("From worker", wn, ":", data)
					time.Sleep(1000 * time.Millisecond)
				default:
				}

				// неблокирующая проверка контекста на завершение
				select {
				case <-ctx.Done():
					fmt.Println("Worker", wn, "stopped")
					wg.Done()
					return
				default:
				}
			}

		}(i, ctx, &wg)
	}

	// запуск горутины для записи данных в канал
	go func(ctx context.Context, wg *sync.WaitGroup) {
		for i := 0; ; i++ {
			// запись данных в канал
			channel <- i
			time.Sleep(200 * time.Millisecond)

			// неблокирующая проверка контекста на завершение
			select {
			case <-ctx.Done():
				fmt.Println("Sender stopped")
				return
			default:
			}
		}

	}(ctx, &wg)

	signalCh := make(chan os.Signal, 1)   // создание канала для отслеживание сигнала о прерывании
	signal.Notify(signalCh, os.Interrupt) // направление сигнала о прерывании в канал

	// запуск горутины для завершения контекста при получении сигнала о прерывании
	go func() {
		select {
		case <-signalCh:
			cancel() // завершение контекста
			return
		}
	}()

	wg.Wait() // ожидание завершения всех горутин
	fmt.Println("All goroutines stopped, exiting")

	fmt.Println("Remaining data in channel:")
	flag := false
	for {
		if flag {
			break
		}
		// неблокирующее получение данных из канала
		select {
		case data := <-channel:
			fmt.Printf("%d\n", data)
		default:
			fmt.Printf("That's all")
			flag = true
		}
	}
}
