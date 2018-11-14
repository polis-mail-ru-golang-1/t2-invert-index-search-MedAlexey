package main

import (
	"fmt"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/config"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/invertIndex"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/web"
	"github.com/rs/zerolog"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	invertIndexMap = make(invertIndex.Index)
	sliceOfNames   []string
	portNumber     int
)

func main() {

	logFile, err := os.OpenFile("logFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening logFile: %v", err)
		os.Exit(1)
	}
	defer logFile.Close()
	mainLogger := zerolog.New(logFile).With().Timestamp().Logger()
	web.PageLogger = mainLogger

	arg := os.Args[1:]

	if len(arg) < 1 {
		fmt.Println("No arguments.")
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	config.Load(arg[0])
	configuration := config.Configuration

	portNumber = getPortNumberFromString(configuration.Listen)
	sliceOfNames = getFilesFromDir(configuration.Dir)
	directoryOfFiles := configuration.Dir

	mainLogger.Info().Msgf("Server starting at :%s", strconv.Itoa(portNumber))

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
		fmt.Println("Wrong listen nubmer: " + "\"" + arg + "\"")
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
