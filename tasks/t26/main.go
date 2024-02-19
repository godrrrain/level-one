package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
*/

func isUnique(str string) bool {
	str = strings.ToLower(str)
	// пустая мапа для хранения уникальных символов
	uniqueChars := make(map[rune]bool)

	for _, char := range str {
		// проверяем, есть ли этот символ уже в мапе
		if uniqueChars[char] {
			return false
		}
		uniqueChars[char] = true
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	str4 := "abA"

	fmt.Println(isUnique(str1))
	fmt.Println(isUnique(str2))
	fmt.Println(isUnique(str3))
	fmt.Println(isUnique(str4))
}
