package makeInvertIndex

import (
	"strings"
	"sync"
)

type Index map[string]map[string]int

//построение обратного индекса файла
func MakeInvertIndexForFile(file string, fileName string) Index {

	invertIndexMap := make(Index)
	words := strings.Split(string(file), " ")

	for _, word := range words {

		word = trimWord(word)

		if word != "" {
			addWordToMap(word, invertIndexMap, fileName)
		}
	}

	return invertIndexMap
}

func trimWord(word string) string {

	word = strings.ToLower(word)
	word = strings.TrimFunc(word, func(r rune) bool {
		return (r >= 0 && r <= 64) || (r >= 91 && r <= 96) || (r >= 123)
	})

	return word
}

//добавляем слово в map
func addWordToMap(word string, invertIndexMap Index, fileName string) {

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

//добавление индекса одного файла в общий индекс
func AddFileIndexToMain(mainIndex Index, fileIndex Index, mutex *sync.RWMutex) {

	for key, value := range fileIndex {
		mutex.RLock()
		if _, ok := mainIndex[key]; !ok {
			mutex.RUnlock()
			mutex.Lock()
			mainIndex[key] = value
			mutex.Unlock()
		} else {
			for key1, value1 := range fileIndex[key] {
				mutex.RUnlock()
				mutex.Lock()
				mainIndex[key][key1] = value1
				mutex.Unlock()
			}
		}
	}
}
