package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReplaceWord(word, newWord, path string) {
	src := "modified.txt"
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		fmt.Println("To count words enter  <path> <searchedWord> or <word> <newWord> <path> if you want to replace some word")
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(word), []byte(newWord), -1)
	if err = ioutil.WriteFile(src, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.Rename(src, path) //!!
	if err != nil {
		return
	}
}

func WordCount(path, searchedWord string) int {
	counts := map[string]int{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //!

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()

		word = strings.ToLower(word)
		counts[word]++
		// fmt.Println(counts)
	}
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Number of count is", counts[strings.ToLower(searchedWord)])
	return counts[strings.ToLower(searchedWord)]
}

func main() {
	params := os.Args

	if len(params) == 3 {
		_, err := fmt.Println(WordCount(params[1], params[2]))
		if err != nil {
			return
		}
		return
	} else if len(params) == 4 {
		ReplaceWord(params[1], params[2], params[3])
		return
	}

	fmt.Println("To count words enter  <path> <searchedWord> or <word> <newWord> <path> if you want to replace some word")
}
