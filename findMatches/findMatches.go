package findMatches

import (
	"strconv"
	"strings"
	"sync"
)

func FindMatches(phrase string, invertIndexMap map[string]map[string]int, fileNames []string) ([][]string, [][]string) {

	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	// слайс файлов, в которых фраза существует полностью
	var fullMatches [][]string
	//слайс файлов, в которых фраза существует не полностью
	var notFullMatches [][]string

	sPhrase := strings.Split(phrase, " ")

	for _, fileName := range fileNames {

		wg.Add(1)
		go func(fullMatches *[][]string, notFullMatches *[][]string, sPhrase []string, fileName string, wg *sync.WaitGroup) {

			defer wg.Done()
			fullPhrase := true
			numbOfMatches := 0

			for _, word := range sPhrase {

				word = strings.TrimSpace(word)

				if word != "" {
					if _, ok := invertIndexMap[word][fileName]; !ok {
						fullPhrase = false
					} else {
						numbOfMatches += invertIndexMap[word][fileName]
					}
				}
			}

			if numbOfMatches != 0 {
				addedFile := make([]string, 0)
				addedFile = append(addedFile, fileName)
				addedFile = append(addedFile, strconv.Itoa(numbOfMatches))
				if fullPhrase {
					mutex.Lock()
					*fullMatches = append(*fullMatches, addedFile)
					*fullMatches = shiftLastElement(*fullMatches)
					mutex.Unlock()
				} else {
					mutex.Lock()
					*notFullMatches = append(*notFullMatches, addedFile)
					*notFullMatches = shiftLastElement(*notFullMatches)
					mutex.Unlock()
				}
			}
		}(&fullMatches, &notFullMatches, sPhrase, fileName, wg)
	}

	wg.Wait()
	return fullMatches, notFullMatches
}

// перемещаем новый (последний) элемент слайса ближе к началу
// ( сортируем элементы по числу совпадений: от большего к меньшему)
func shiftLastElement(matches [][]string) [][]string {

	for i := len(matches) - 1; i > 0; i-- {
		curElementMatches, _ := strconv.Atoi(matches[i][1])
		nextElementMatches, _ := strconv.Atoi(matches[i-1][1])
		if curElementMatches > nextElementMatches {
			matches[i], matches[i-1] = matches[i-1], matches[i]
		}
	}

	return matches
}
