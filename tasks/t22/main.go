package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20
*/

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("48230439432383944", 10)
	b.SetString("12385433298091834", 10)

	// Перемножение
	mul := new(big.Int)
	mul.Mul(a, b)
	fmt.Println("Умножение:", mul)

	// Деление
	div := new(big.Int)
	div.Div(a, b)
	fmt.Println("Деление:", div)

	// Сложение
	add := new(big.Int)
	add.Add(a, b)
	fmt.Println("Сложение:", add)

	// Вычитание
	sub := new(big.Int)
	sub.Sub(a, b)
	fmt.Println("Вычитание:", sub)
}
