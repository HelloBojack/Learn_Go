package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()
	switch {
	case hour > 21:
		fmt.Println("下班啦")
	default:
		fmt.Println("没下班")
	}
}
