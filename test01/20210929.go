package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("小额代付", "支付"))
	fmt.Println(strings.Contains("代付", "小额代付"))
}
