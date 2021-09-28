package main

import (
	"fmt"
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
	// 计算
	number1 := "2"
	number2 := "3"
	println("Sum:", sum(number1, number2))
	// defer 关键字
	name := "go"
	defer printName(name) // output: go
	name = "python"
	defer printName(name) // output: python
	name = "java"
	printName(name) // output: java
	deferTestA()
}

func deferTestA() {
	i := 0

	defer fmt.Println(i) //因为i=0，所以此时就明确告诉golang在程序退出时，执行输出0的操作
	i++
	defer fmt.Println(i) //输出1，因为i此时就是1
	return
}

func printName(name string) {
	fmt.Println(name)
}

func sum(a string, b string) int {
	int_a, _ := strconv.Atoi(a)
	int_b, _ := strconv.Atoi(b)
	return int_a + int_b
}
