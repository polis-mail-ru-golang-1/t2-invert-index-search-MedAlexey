package main

import (
	"fmt"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/invertIndex"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/web"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	invertIndexMap = make(map[string]map[string]int)
	sliceOfNames   []string
	portNumber     int
)

func main() {

	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	arg := os.Args[1:]

	if len(arg) != 2 {
		fmt.Println("Wrong number of arguments.")
		os.Exit(1)
	}

	portNumber = getPortNumberFromArg(arg[1])
	sliceOfNames = getFilesFromArg(arg[0])

	for _, fileName := range sliceOfNames {

		wg.Add(1)
		go func(invertIndexMap map[string]map[string]int, fileName string, wg *sync.WaitGroup) {
			defer wg.Done()

			file, err := ioutil.ReadFile(arg[0] + "/" + fileName)
			if err != nil {
				fmt.Print("Error opening file " + "\"" + arg[0] + fileName + "\"\n")
				os.Exit(1)
			}

			indexMap := invertIndex.MakeInvertIndexForFile(string(file), fileName)

			invertIndex.AddFileIndexToMain(invertIndexMap, indexMap, mutex)

		}(invertIndexMap, fileName, wg)

	}

	wg.Wait()

	web.InvertIndexMap = invertIndexMap
	web.SliceOfNames = sliceOfNames

	http.HandleFunc("/", web.SearchPage)
	http.HandleFunc("/result", web.ResultPage)
	http.ListenAndServe(":"+strconv.Itoa(portNumber), nil)
}

func getPortNumberFromArg(arg string) int {

	arg = strings.TrimLeft(arg, ":")
	number, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Wrong port nubmer: " + "\"" + arg + "\"")
		os.Exit(2)
	}

	return number
}

// выделение имён файлов в заданной директории в слайс
func getFilesFromArg(dir string) []string {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	sliceOfNames := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() {
			sliceOfNames = append(sliceOfNames, file.Name())
		}
	}

	return sliceOfNames
}
