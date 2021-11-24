package main

import (
	"fmt"
	"time"
)

// 微信公众号:技术生活
// _switch_ ，方便的条件分支语句。
func main() {
	// 一个基本的 switch
	i := 2
	fmt.Println("i=", i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	// 在一个 case 语句中，你可以使用逗号来分割多个表达式。

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")

	}
	// 不带表达式的 switch 是实现 if/else 的一种简单方式。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")

	}
}
