package main

import (
	"bufio"
	"fmt"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/findMatches"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/makeInvertIndex"
	"os"
)

func main() {

	arg := os.Args[1:]

	if len(arg) < 2 {
		fmt.Println("No arguments.")
		os.Exit(1)
	}

	invertIndexMap := make(map[string]map[string]int)
	for _, fileName := range arg {
		makeInvertIndex.MakeInvertIndexForFile(fileName, invertIndexMap)
	}

	var phrase string
	fmt.Println("Enter your phrase:")
	phrase = scan()
	fullMatches, notFullMatches := findMatches.FindMatches(phrase, invertIndexMap, arg)

	printMatches(fullMatches, "Файлы, в которых фраза присутствует полностью:")
	printMatches(notFullMatches, "Файлы, в которых фраза присутствует не полностью:")
}

func printMatches(matches [][]string, message string) {

	if len(matches) != 0 {
		fmt.Println(message)
		for _, file := range matches {
			fmt.Println("-", file[0], ";", "совпадений -", file[1])
		}
	}
}

// чтение фразы из stdin
func scan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return str
}
