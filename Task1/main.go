package main

import (
	"fmt"
	"strings"
)

func draw(x, y int) {
	symbol := "*"
	var lineArr []string
	for i := 0; i < x; i++ {
		lineArr = append(lineArr, symbol)
	}
	for i := 0; i < y; i++ {
		if i%2 == 0 {
			fmt.Println(strings.Join(lineArr, " "))
		} else {
			fmt.Println(" ", strings.Join(lineArr, " "))
		}
	}
}

var x int = 5
var y int = 5

func main() {
	draw(x, y)
}
