package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func Test_findNeededNumbers(t *testing.T) {
	num = 25
	expected := []int{0, 1, 2, 3, 4}
	result := findNeededNumbers()
	fmt.Println(result)
	if !reflect.DeepEqual(expected, result) { //expected != result
		t.Errorf("expected %v, but got %v", expected, result)
	}
}

func Test_dataInput(t *testing.T) {
	osArgsWere := os.Args
	defer func() { os.Args = osArgsWere }()

	testcases := map[bool]string{
		true:  "-n=115",
		false: "-n=-10",
	}

	for key, val := range testcases {
		os.Args[1] = val
		got := dataInput()
		if got != key {
			t.Errorf("Expected %v, but received %v", key, got)
		}
	}
}
