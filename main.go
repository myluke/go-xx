package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	fmt.Println("Hi,Luke")
	var defaultFloat32 float32
	var defaultInt int64
	var defaultString string
	// 变量赋值
	defaultString = "Hi,Luke"
	// 打印
	fmt.Printf("defaultFloat32:%+v,defaultInt:%+v,defaultString:%+v", defaultFloat32, defaultInt, defaultString)
	// 计算
	num1 := "2"
	num2 := "3"
	fmt.Println("sum:", sum(num1, num2))
	var blockType string
	fmt.Println(len(blockType))
	// 时间和日期操作
	t := time.Now().UTC()
	fmt.Println(t)
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println("休息2秒")
	time.Sleep(2 * time.Second)
	t = time.Now().UTC()
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// 增加30天
	m := time.Now().AddDate(0, 0, 30)
	fmt.Printf("m:%s", m)

}

// 定义函数
func sum(a string, b string) int {
	// 字符串转整型
	intA, _ := strconv.Atoi(a)
	intB, _ := strconv.Atoi(b)
	return intA + intB
}
