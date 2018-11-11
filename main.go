package main

import (
	"fmt"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/config"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/invertIndex"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/logger"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/web"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	invertIndexMap = make(invertIndex.Index)
	sliceOfNames   []string
	portNumber     int
)

func main() {

	logger.StartTime = time.Now()

	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	configuration := config.Load()

	portNumber = getPortNumberFromString(configuration.Port)
	sliceOfNames = getFilesFromDir(configuration.Dir)
	directoryOfFiles := configuration.Dir

	logger.PrintLog("Starting server at " + strconv.Itoa(portNumber))

	for _, fileName := range sliceOfNames {

		wg.Add(1)
		go func(invertIndexMap invertIndex.Index, fileName string, wg *sync.WaitGroup) {
			defer wg.Done()

			file, err := ioutil.ReadFile(directoryOfFiles + "/" + fileName)
			if err != nil {
				fmt.Print("Error opening file " + "\"" + directoryOfFiles + fileName + "\"\n")
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

func getPortNumberFromString(arg string) int {

	arg = strings.TrimLeft(arg, ":")
	number, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Wrong port nubmer: " + "\"" + arg + "\"")
		os.Exit(2)
	}

	return number
}

// выделение имён файлов в заданной директории в слайс
func getFilesFromDir(dir string) []string {

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
