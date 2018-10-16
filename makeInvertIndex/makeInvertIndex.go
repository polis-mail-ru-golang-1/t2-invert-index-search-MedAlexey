package makeInvertIndex

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 13 - перенос строки; 32 - space
//построение обратного индекса файла
func MakeInvertIndexForFile(fileName string, invertIndexMap map[string][]string) {

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
}

//добавляем слово в map
func addWordToMap(word string, invertIndexMap map[string][]string, fileName string) {

	//если ключ не существует, создаём его
	if _, ok := invertIndexMap[word]; !ok {
		invertIndexMap[word] = make([]string, 0)
	}

	invertIndexMap[word] = append(invertIndexMap[word], fileName)

}
