package main

import (
	"bufio"
	"fmt"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/findMatches"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/makeInvertIndex"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func main() {

	invertIndexMap := make(map[string]map[string]int)
	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	arg := os.Args[1:]

	if len(arg) < 2 {
		fmt.Println("No arguments.")
		os.Exit(1)
	}

	for _, fileName := range arg {

		wg.Add(1)
		go func(invertIndexMap map[string]map[string]int, fileName string, wg *sync.WaitGroup) {
			defer wg.Done()

			file, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Print("Error opening file " + "\"" + fileName + "\"\n")
				os.Exit(1)
			}

			//TODO переводить в нижний регистр и убирать с боков всё, что не буква
			sFile := strings.Split(string(file), " ")

			indexMap := makeInvertIndex.MakeInvertIndexForFile(sFile, fileName)

			makeInvertIndex.AddFileIndexToMain(invertIndexMap, indexMap, mutex)

		}(invertIndexMap, fileName, wg)
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
