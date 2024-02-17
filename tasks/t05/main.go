package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.

*/

func main() {
	var t int
	fmt.Println("Введи время в секундах")
	fmt.Scan(&t)

	c := make(chan int)

	go writer(c)
	go reader(c)
	time.Sleep(time.Duration(t) * time.Second)
}

func writer(c chan int) {
	for i := 0; ; i++ {
		c <- i
		time.Sleep(time.Millisecond * 100)
	}
}

func reader(c chan int) {
	for val := range c {
		fmt.Println(val)
	}
}
