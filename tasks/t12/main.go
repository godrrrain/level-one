package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	dataMap := make(map[string]int)
	for index, value := range data {
		dataMap[value] = index
	}
	fmt.Println(dataMap)
}
