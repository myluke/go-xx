package main

import "fmt"

func main() {
	// 最常用的方式，带单个循环条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// 经典的初始化/条件/后续形式 `for` 循环。
	for j := 7; j <= 10; j++ {
		fmt.Println(j)
	}
	// 不带条件的 ·for· 循环将一直执行，直到在循环体内使用了break 或 return
	for {
		fmt.Println("loop")
		break
	}

	//
}
