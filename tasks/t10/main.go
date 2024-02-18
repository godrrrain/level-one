package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.
*/

func main() {
	groupData := make(map[int][]float64)

	tempArr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, temp := range tempArr {
		index := (int(temp) / 10) * 10
		groupData[index] = append(groupData[index], temp)
	}

	fmt.Println("Группы")
	for key, group := range groupData {
		fmt.Println(key, ":", group)
	}
}
