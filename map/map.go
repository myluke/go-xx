package main

import "fmt"

// 微信公众号:技术生活
// 视频原创不易，请多多关注！多多点赞。
// _map_ 是 Go 内置[关联数据类型]
// 在一些其他的语言中称为_哈希_ 或者_字典_ ）

func main() {
	// 使用典型的 ·make[key]= val` 语法来设置键值对
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// 使用 name[key] 来获取值
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	// map 调用 内建的 len 来获取键值对的数量
	fmt.Println("len:", len(m))
	// 使用 delete 删除一个键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// 你也可以通过这个语法在同一行申明和初始化一个新的map

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}

// 微信公众号:技术生活
// 视频原创不易，请多多关注！多多点赞。
