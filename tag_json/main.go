package main

import (
	"encoding/json"
	"fmt"
)

type Move struct {
	Actor []string `json:"actor"`
	Title string   `json:"title"`
	Year  int      `json:"year"`
}

func main() {
	m1 := Move{[]string{"xlk"}, "xlk", 1234}

	str1, err := json.Marshal(m1)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("%v %s", m1, str1)
}
