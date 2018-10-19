package findMatches

import (
	"strconv"
	"strings"
)

func FindMatches(phrase string, invertIndexMap map[string]map[string]int, fileNames []string) ([][]string, [][]string) {

	// слайс файлов, в которых фраза существует полностью
	var fullMatches [][]string
	//слайс файлов, в которых фраза существует не полностью
	var notFullMaches [][]string

	sPhrase := strings.Split(phrase, " ")

	for _, fileName := range fileNames {
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
				fullMatches = append(fullMatches, addedFile)
				fullMatches = shiftLastElement(fullMatches)
			} else {
				notFullMaches = append(notFullMaches, addedFile)
				notFullMaches = shiftLastElement(notFullMaches)
			}
		}
	}

	return fullMatches, notFullMaches
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
