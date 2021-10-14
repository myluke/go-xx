package main

import (
	"fmt"
	"time"
)

// 循环 5个 等待 1秒
func main() {

	totalUser := 10
	batchSendNumber := 5
	// 遍历所有用户
	for i := 0; i < totalUser; i += batchSendNumber {
		// 每5个一批

		for j := 0; j < batchSendNumber; j++ {
			idx := i + j
			fmt.Print("j:")
			fmt.Println(j)
			fmt.Print("idx:")
			fmt.Println(idx)
		}
		fmt.Println("------------")
		fmt.Print("i:")
		fmt.Println(i)
		// sleep 1s
		time.Sleep(1 * time.Second)
	}

}
