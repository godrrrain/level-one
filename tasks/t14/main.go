package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

func typeOf1(i interface{}) string {
	switch i.(type) {
	case string:
		return "String"
	case int:
		return "Number"
	case bool:
		return "Bool"
	default:
		return "Other"
	}
}

func typeOf2(i interface{}) string {
	return reflect.TypeOf(i).String()
}

func main() {
	var i interface{} = 230
	fmt.Println(typeOf1(i))
	fmt.Println(typeOf2(i))
}
