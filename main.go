package main

import (
	"bufio"
	"fmt"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/findMatches"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/makeInvertIndex"
	"os"
	"sync"
)

func main() {

	invertIndexMap := make(map[string]map[string]int)
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	arg := os.Args[1:]

	if len(arg) < 2 {
		fmt.Println("No arguments.")
		os.Exit(1)
	}

	for _, fileName := range arg {

		wg.Add(1)
		go func(invertIndexMap map[string]map[string]int, fileName string, wg *sync.WaitGroup, mutex *sync.Mutex) {
			defer wg.Done()
			indexMap := makeInvertIndex.MakeInvertIndexForFile(fileName)

			for key, value := range indexMap {
				if _, ok := invertIndexMap[key]; !ok {
					mutex.Lock()
					invertIndexMap[key] = value
					mutex.Unlock()
				} else {
					for key1, value1 := range indexMap[key] {
						mutex.Lock()
						invertIndexMap[key][key1] = value1
						mutex.Unlock()
					}
				}
			}
		}(invertIndexMap, fileName, wg, mutex)
	}

	wg.Wait()

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
