package findMatches

import (
	"strconv"
	"strings"
)

func FindMatches(phrase string, invertIndexMap map[string][]string) [][]string {

	//массив совпадений
	var matches [][]string

	sPhrase := strings.Split(phrase, " ")

	for _, word := range sPhrase {

		if word != "" {
			//удаляем последний символ последнего слова (enter)
			if word[len(word)-1] == 10 {
				word = word[:len(word)-1]
			}

			if _, ok := invertIndexMap[word]; ok {
				fileNames := invertIndexMap[word]

				//добавляем имена файлов в слайс соответствий
				matches = addNames(fileNames, matches)
			}
		}

	}

	return matches
}

func addNames(fileNames []string, matches [][]string) [][]string {

	for _, fileName := range fileNames {
		fileNameExistInMatches := false
	addFileName:
		for i, name := range matches {

			if len(name) > 0 && fileName == name[0] {
				numOfMatches, _ := strconv.Atoi(name[1])
				numOfMatches += 1
				name[1] = strconv.Itoa(numOfMatches)
				shiftElement(&matches, i)
				fileNameExistInMatches = true
				break addFileName
			}
		}

		//если имя файла нет в массиве совпадений, то добавляем его
		if !fileNameExistInMatches {
			tmp := make([]string, 0)
			tmp = append(tmp, fileName)
			tmp = append(tmp, "1")
			matches = append(matches, tmp)
		}
	}

	return matches
}

// перемещаем элемент слайса, в котором увеличилось число совпадений, ближе к началу
// ( сортируем элементы по числу совпадений: от большего к меньшему)
func shiftElement(matches *[][]string, startI int) {
	tmp := (*matches)[startI]
	tmpMatches, _ := strconv.Atoi(tmp[1])

	for i := startI - 1; i > -1; i-- {
		curElementMatches, _ := strconv.Atoi((*matches)[i][1])
		if curElementMatches < tmpMatches {
			(*matches)[i+1] = (*matches)[i]
			(*matches)[i] = tmp
		}
	}

}
