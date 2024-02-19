package main

import (
	"fmt"
	"slices"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func reverseStr1(inp string) string {
	inpWords := strings.Split(inp, " ")

	length := len(inpWords)
	for i := 0; i < (length)/2; i++ {
		inpWords[i], inpWords[length-i-1] = inpWords[length-i-1], inpWords[i]
	}
	return strings.Join(inpWords, " ")
}

func reverseStr2(inp string) string {
	inpWords := strings.Split(inp, " ")
	slices.Reverse(inpWords)
	return strings.Join(inpWords, " ")
}

func main() {
	input := "snow dog sun"
	fmt.Println("Исходная строка:", input)

	res1 := reverseStr1(input)
	fmt.Println("Используя простую перестановку:", res1)

	res2 := reverseStr2(input)
	fmt.Println("Используя slices.Reverse():", res2)
}
