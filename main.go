package main

import (
	"bufio"
	"fmt"
	"os"
	"t2-invert-index-search-MedAlexey/findMatches"
	"t2-invert-index-search-MedAlexey/makeInvertIndex"
)

func main() {
	invertIndexMap := make(map[string][]string)
	arg := os.Args[1:]

	if len(arg) < 2 {
		fmt.Println("Enter files as arguments.")
		os.Exit(1)
	}

	//создаём инвертированный индекс для файлов
	for _, fileName := range arg {
		makeInvertIndex.MakeInvertIndexForFile(fileName, invertIndexMap)
	}

	//читаем поисковую фразу
	var phrase string
	fmt.Println("Enter your phrase.")
	phrase = scan()

	// массив совпадений
	matches := make([][]string, 0)

	//ищем соответствия
	matches = findMatches.FindMatches(phrase, invertIndexMap, matches)

	printMatches(matches)
}

func printMatches(matches [][]string) {
	for i := range matches {
		fmt.Println("-", matches[i][0], ";", "совпадений -", matches[i][1])
	}
}

func scan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return str
}
