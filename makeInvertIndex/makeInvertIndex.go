package makeInvertIndex

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 13 - перенос строки; 32 - space
//построение обратного индекса файла
func MakeInvertIndexForFile(fileName string, invertIndexMap map[string][]string) {

	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Print("Error opening file " + "\"" + fileName + "\"\n")
		os.Exit(1)
	}

	sFile := string(file)

	var word string
mapCreating:
	for true {
		// берём очередное слово файла
		word, sFile = NextWord(sFile)

		if word != "" {
			addWordToMap(word, invertIndexMap, fileName)
		}

		if sFile == "" {
			break mapCreating
		}
	}
}

//добавляем слово в map
func addWordToMap(word string, invertIndexMap map[string][]string, fileName string) {

	//если ключ не существует, создаём его
	if _, ok := invertIndexMap[word]; !ok {
		invertIndexMap[word] = make([]string, 0)
	}

	invertIndexMap[word] = append(invertIndexMap[word], fileName)

}

// выделяем из файла первое слово и обрезаем файл
func NextWord(file string) (string, string) {

	// попускаем пробелы и переносы строк
removeGaps:
	for _, letter := range file {
		if letter == 13 || letter == 32 {
			file = file[1:]
		} else {
			break removeGaps
		}
	}

	result := ""

createWord:
	for i, letter := range file {
		if letter == 13 /* перенос строки*/ || letter == 32 /*пробел*/ {
			break createWord
		}
		result = file[:i+1]
	}
	if len(file) > len(result) {
		file = file[len(result):]
	} else {
		file = ""
	}

	return result, file
}
