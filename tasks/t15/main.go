package main

import (
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}
*/

func createHugeString(a int) string {
	str := "0"
	res := strings.Repeat(str, a)
	return res
}

func someFunc() {
	v := []rune(createHugeString(1 << 10))
	runeArr := make([]rune, 100)

	copy(runeArr, v) //чтобы не использовать ссылку на большую строку
	res := string(runeArr)
	fmt.Println("Размер:", len(runeArr), "Значение:", res)
}

func main() {
	someFunc()
}
