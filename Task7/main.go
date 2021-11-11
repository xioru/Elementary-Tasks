package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	num int
)

const (
	delimiter    = ", "
	Instructions = `(You should specify natural number(>=0) to analyze by passing it to function call)
	(task7 -n 225	Will print 'Result is: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14`)

func dataInput() bool {
	flag.Parse()
	successfulInput := (num >= 0)
	return successfulInput
}

func findNeededNumbers() []int {
	bound := int(math.Sqrt(float64(num)))
	if bound*bound >= num {
		bound--
	}
	result := make([]int, bound+1)
	for i := 0; i <= bound; i++ {
		result[i] = i
	}
	return result
}

func init() {
	flag.IntVar(&num, "n", 0, "Number to analyze")
}

func main() {
	if dataInput() {
		fmt.Print("Result is: ")
		for ind, res := range findNeededNumbers() {
			if ind > 0 {
				fmt.Print(delimiter)
			}
			fmt.Printf("%d", res)
		}
		fmt.Println()

	} else {
		fmt.Println(Instructions)
	}
}
