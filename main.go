package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Hi, Luke")

	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)

	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	println(i, s)

	number1 := os.Args[1]
	number2 := os.Args[2]
	println("Sum:", sum(number1, number2))
}

func sum(a string, b string) int {
	int_a, _ := strconv.Atoi(a)
	int_b, _ := strconv.Atoi(b)
	return int_a + int_b
}
