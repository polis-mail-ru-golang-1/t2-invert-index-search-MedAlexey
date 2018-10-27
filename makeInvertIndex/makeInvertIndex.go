package makeInvertIndex

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//построение обратного индекса файла
func MakeInvertIndexForFile(fileName string) map[string]map[string]int {
	invertIndexMap := make(map[string]map[string]int)
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Print("Error opening file " + "\"" + fileName + "\"\n")
		os.Exit(1)
	}

	//разбиваем файл на слова и добавляем каждое слово в map
	sFile := strings.Split(string(file), " ")

	for _, word := range sFile {

		if word != "" {
			addWordToMap(word, invertIndexMap, fileName)
		}
	}

	return invertIndexMap
}

//добавляем слово в map
func addWordToMap(word string, invertIndexMap map[string]map[string]int, fileName string) {

	//если ключ не существует, создаём его
	if _, ok := invertIndexMap[word]; !ok {
		newFile := make(map[string]int)
		newFile[fileName] = 1
		invertIndexMap[word] = newFile
	} else {

		if _, ok := invertIndexMap[word][fileName]; !ok {
			invertIndexMap[word][fileName] = 1
		} else {
			invertIndexMap[word][fileName]++
		}
	}
}
