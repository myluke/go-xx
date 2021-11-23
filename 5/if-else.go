package main

import "fmt"

// `if` 和 `else` 分支结构在 Go 中当然是直接了当的了。
func main() {
	// 这里是一个基本的例子
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	// 你可以不要 else 只用 if 语句
	if 8%4 == 0 {
		fmt.Println("8 is xxx")
	}
	// 在条件语句之前可以有一个语句
	if num := 9; num < 0 {
		fmt.Println(" is negative")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}
}
