package main

import (
	"fmt"
	"slices"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func reverseStr1(inp string) string {
	inpRunes := []rune(inp)

	length := len(inpRunes)
	for i := 0; i < (length)/2; i++ {
		inpRunes[i], inpRunes[length-i-1] = inpRunes[length-i-1], inpRunes[i]
	}
	return string(inpRunes)
}

func reverseStr2(inp string) string {
	inpRunes := []rune(inp)
	slices.Reverse(inpRunes)
	return string(inpRunes)
}

func main() {
	input := "главрыба"
	fmt.Println("Исходная строка:", input)

	res1 := reverseStr1(input)
	fmt.Println("Используя простую перестановку:", res1)

	res2 := reverseStr2(input)
	fmt.Println("Используя slices.Reverse():", res2)
}
