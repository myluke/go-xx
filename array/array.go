package main

import "fmt"

// 微信公众号： 技术生活
// 在 Go 中，_数组_ 是一个固定长度的数列。
func main() {
	// 创建一个数组，长度为 5
	var a [5]int
	fmt.Println("emp:", a)
	// 我们可以使用 array[index] = value 的方式来设置数组元素的值
	// 指定位置的值，或者用 array[index] 取出指定位置的值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	// 使用 `len` 函数获取数组的长度
	fmt.Println("len:", len(a))
	// 使用 `cap` 函数获取数组的容量
	fmt.Println("cap:", cap(a))
	// 使用 这个语法在一行中声明并初始化一个数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	// 数组的存储类型是单一的，但是你可以组合这些数据
	// 来构造多维的数组结构
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

// 微信公众号： 技术生活
