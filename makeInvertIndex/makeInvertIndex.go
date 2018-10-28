package makeInvertIndex

import "sync"

type Index map[string]map[string]int

//построение обратного индекса файла
func MakeInvertIndexForFile(words []string, fileName string) Index {

	invertIndexMap := make(Index)

	for _, word := range words {

		if word != "" {
			addWordToMap(word, invertIndexMap, fileName)
		}
	}

	return invertIndexMap
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
