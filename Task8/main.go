package main

import (
	"flag"
	"fmt"
)

var (
	num1, num2 int
)

const (
	delimiter = ", "
)

func main() {
	if dataInput() {
		printFibonacciSeries(num2)
		// fmt.Print("Result is: ")
		// for ind, res := range printFibonacciSeries(num2) {
		// 	if ind > 0 {
		// 		fmt.Print(delimiter)
		// 	}
		// 	fmt.Printf("%d", res)
		// }
		// fmt.Println()

	} else {
		printInstructions()
	}
}

func printFibonacciSeries(num int) {
	a := 0
	b := 1
	c := b
	fmt.Printf("Series is: %d%s%d", a, delimiter, b)
	for {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf("%s%d", delimiter, b)
	}
}

func dataInput() bool {
	flag.IntVar(&num1, "n1", 0, "First number: ")
	flag.IntVar(&num2, "n2", 0, "Second number: ")
	flag.Parse()
	successfulInput := (num2 >= 0)
	return successfulInput
}

func printInstructions() {
	fmt.Println("You should specify natural number(>=0) to analyze by passing it to function call")
	fmt.Println("\ttask7 -n 225\t//will print 'Result is: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14'")
}
