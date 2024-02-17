package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины
*/

func testGoroutine1(quit chan bool, wg *sync.WaitGroup) {
	for {
		// неблокирующее получение данных из канала
		select {
		case <-quit:
			fmt.Println("[1] Stopping the goroutine")
			wg.Done()
			return
		default:
			fmt.Println("[1] Goroutine is still running")
			time.Sleep(time.Second)
		}
	}
}

func testGoroutine2(stopFlag *bool, wg *sync.WaitGroup) {

	for {
		// постоянная проверка логической переменной по ссылке
		if *stopFlag {
			fmt.Println("[2] Stopping the goroutine")
			wg.Done()
			return
		}
		fmt.Println("[2] Goroutine is still running")
		time.Sleep(time.Second)
	}
}

func testGoroutine3(ctx context.Context, wg *sync.WaitGroup) {
	for {
		// неблокирующая проверка на завершение контекста
		select {
		case <-ctx.Done():
			fmt.Println("[3] Stopping the goroutine")
			wg.Done()
			return
		default:
			fmt.Println("[3] Goroutine is still running")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	//вызываем первую горутину
	quit := make(chan bool)
	wg.Add(1)
	go testGoroutine1(quit, &wg)

	//вызываем вторую горутину
	var stopFlag bool
	wg.Add(1)
	go testGoroutine2(&stopFlag, &wg)

	//вызываем третью горутину
	ctx, cancel := context.WithCancel(context.Background()) // создание контекста для контроля работы горутины
	wg.Add(1)
	go testGoroutine3(ctx, &wg)

	//вырубаем первую горутину
	time.Sleep(time.Second)
	quit <- true

	//вырубаем вторую горутину
	time.Sleep(time.Second)
	stopFlag = true

	//вырубаем третью горутину
	time.Sleep(time.Second)
	cancel()

	wg.Wait()
}
