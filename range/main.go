package main

import "fmt"

// 微信公众号:技术生活
// 视频原创不易，请多多关注！多多点赞。
// range 用于迭代各种各样的数据结构。
// 让我们来看看如何在我们已经学过的数据结构上使用 range。

func main() {
	// 求和
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// 遍历数组，取需要的索引
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range 在map中迭代键值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s|", k, v)
	}
	// range 在map中迭代键
	for k := range kvs {
		fmt.Println("key:", k)
	}
	// range 在字符串中迭代unicode编码
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

// 微信公众号:技术生活
// 视频原创不易，请多多关注！多多点赞。
