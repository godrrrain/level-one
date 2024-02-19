package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/

// структура Point
type Point struct {
	x, y float64
}

// конструктор Point
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// метод для вычисления расстояния между двумя точками
func (p *Point) GetDistanceTo(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(5, 3)
	p2 := NewPoint(3, 6)

	// вычисление расстояния между точками
	distance := p1.GetDistanceTo(p2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
