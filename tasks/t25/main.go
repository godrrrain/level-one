package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func sleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}

func main() {
	startTime := time.Now()
	sleep(2)
	fmt.Println("Время работы программы:", time.Since(startTime))
}
