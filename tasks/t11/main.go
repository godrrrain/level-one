package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {

	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	intersection := make(map[int]bool)

	for k := range set1 {
		if set2[k] {
			intersection[k] = true
		}
	}
	fmt.Println(intersection)
}
