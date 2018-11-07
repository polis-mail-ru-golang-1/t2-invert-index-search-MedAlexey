package makeInvertIndex

import (
	"sync"
	"testing"
)

func TestTrimWord(t *testing.T) {

	for i := 0; i < 256; i++ {
		if (i >= 0 && i <= 64) || (i >= 91 && i <= 96) || (i >= 123) {
			inputWord := string(i) + "word" + string(i)
			if result := trimWord(inputWord); result != "word" {
				t.Errorf("trimWord(%q) = %v", inputWord, result)
			}
		}
	}
}

func TestMakeInvertIndexForFile(t *testing.T) {

	var expectedMap = Index{
		"i":       {"file": 2},
		"did":     {"file": 1},
		"not":     {"file": 2},
		"reply":   {"file": 1},
		"to":      {"file": 1},
		"this":    {"file": 1},
		"letter":  {"file": 1},
		"because": {"file": 1},
		"was":     {"file": 1},
		"able":    {"file": 1},
	}

	var inputStr = "I did not reply to this letter, because I was not able"

	if result := MakeInvertIndexForFile(inputStr, "file"); !equalsMaps(result, expectedMap) {
		t.Errorf("trimWord(%q , file) = %v", inputStr, result)
	}

	//-------------------------------------------------------------

	inputStr = "I -did* !/not. ,reply7\" №^to& (this*; %%letter,: ?$because# @I!* (~was` №not:) ?able)"

	if result := MakeInvertIndexForFile(inputStr, "file"); !equalsMaps(result, expectedMap) {
		t.Errorf("trimWord(%q , file) = %v", inputStr, result)
	}
}

func TestAddWordToMap(t *testing.T) {

	var newElements = []struct {
		word     string
		fileName string
	}{
		{"flag", "noon"},
		{"car", "time"},
		{"car", "noon"},
		{"car", "time"},
	}

	var indexMap = Index{}

	// при отсутствии слова
	addWordToMap(newElements[0].word, indexMap, newElements[0].fileName)
	if (len(indexMap) != 1) || (indexMap["flag"]["noon"] != 1) {
		t.Errorf("addWordToMap(%q, %q, %q)", newElements[0].word, indexMap, newElements[0].fileName)
	}
	addWordToMap(newElements[1].word, indexMap, newElements[1].fileName)
	if (len(indexMap) != 2) || (indexMap["car"]["time"] != 1) || (indexMap["flag"]["noon"] != 1) {
		t.Errorf("addWordToMap(%q, %q, %q)", newElements[0].word, indexMap, newElements[0].fileName)
	}

	// при существующем слове и отсутствующем имени файла
	addWordToMap(newElements[2].word, indexMap, newElements[2].fileName)
	if (len(indexMap) != 2) || (indexMap["car"]["time"] != 1) || (indexMap["flag"]["noon"] != 1) ||
		(indexMap["car"]["noon"] != 1) {
		t.Errorf("addWordToMap(%q, %q, %q)", newElements[0].word, indexMap, newElements[0].fileName)
	}

	// при существующем слове и имени файла
	addWordToMap(newElements[3].word, indexMap, newElements[3].fileName)
	if (len(indexMap) != 2) || (indexMap["car"]["time"] != 2) || (indexMap["flag"]["noon"] != 1) ||
		(indexMap["car"]["noon"] != 1) {
		t.Errorf("addWordToMap(%q, %q, %q)", newElements[0].word, indexMap, newElements[0].fileName)
	}
}

func TestAddFileIndexToMain(t *testing.T) {

	mutex := &sync.RWMutex{}

	var expectedMap = Index{
		"flag": {"time": 10, "noon": 5},
		"car":  {"hard": 3},
		"dog":  {"noon": 1},
	}

	var mainMap = Index{
		"flag": {"time": 10},
		"car":  {"hard": 3},
	}

	var addedMap = Index{
		"flag": {"noon": 5},
		"dog":  {"noon": 1},
	}

	if AddFileIndexToMain(mainMap, addedMap, mutex); !equalsMaps(mainMap, expectedMap) {
		t.Errorf("mainMap: %q\n addedMap: %q", mainMap, addedMap)
	}
}

func equalsMaps(curMap Index, expectedMap Index) bool {

	if (curMap == nil) != (expectedMap == nil) {
		return false
	}

	for word, nestedMap := range expectedMap {
		for file := range nestedMap {
			if curMap[word][file] != expectedMap[word][file] {
				return false
			}
		}
	}

	return true
}
