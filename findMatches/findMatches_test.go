package findMatches

import (
	"testing"
)

type Index = map[string]map[string]int

func TestFindMatches(t *testing.T) {

	var fullMatches [][]string
	var notFullMatches [][]string
	var phrase string
	var files = []string{"moon", "file", "star"}

	var indexMap = Index{
		"success":    {"moon": 1},
		"consists":   {"file": 1},
		"of":         {"moon": 2, "star": 1},
		"going":      {"file": 1},
		"from":       {"moon": 1},
		"failure":    {"file": 2, "star": 3},
		"to":         {"moon": 1},
		"without":    {"file": 1},
		"loss":       {"moon": 1},
		"enthusiasm": {"file": 1},
	}

	// фразы нет в файлах
	phrase = "car dog man"
	if resFull, resNotFull := FindMatches(phrase, indexMap, files); !equalsSlices(resFull, fullMatches) || !equalsSlices(resNotFull, notFullMatches) {
		t.Errorf("shiftLastElement(%q, %q, %q) = %q, %q", phrase, indexMap, files, resFull, resNotFull)
	}

	// фраза в файле присутствует не полностью
	phrase = "success going loss"
	notFullMatches = [][]string{
		{"moon", "2"},
		{"file", "1"},
	}

	if resFull, resNotFull := FindMatches(phrase, indexMap, files); !equalsSlices(resFull, fullMatches) || !equalsSlices(resNotFull, notFullMatches) {
		t.Errorf("shiftLastElement(%q, %q, %q) = %q, %q", phrase, indexMap, files, resFull, resNotFull)
	}

	// фраза полностью присутствует в каком-то файле
	phrase = "without enthusiasm"
	fullMatches = [][]string{
		{"file", "2"},
	}
	notFullMatches = nil

	if resFull, resNotFull := FindMatches(phrase, indexMap, files); !equalsSlices(resFull, fullMatches) || !equalsSlices(resNotFull, notFullMatches) {
		t.Errorf("shiftLastElement(%q, %q, %q) = %q, %q", phrase, indexMap, files, resFull, resNotFull)
	}
}

func TestShiftLastElement(t *testing.T) {

	var tests = []struct {
		input  [][]string
		output [][]string
	}{
		//сдвиг до конца
		{[][]string{{"word", "1"}, {"word2", "2"}},
			[][]string{{"word2", "2"}, {"word", "1"}}},

		//сдвиг не до конца
		{[][]string{{"flag", "5"}, {"game", "1"}, {"drive", "4"}},
			[][]string{{"flag", "5"}, {"drive", "4"}, {"game", "1"}}},

		// сдвиг не требуется
		{[][]string{{"dog", "7"}, {"cat", "4"}, {"rat", "1"}},
			[][]string{{"dog", "7"}, {"cat", "4"}, {"rat", "1"}}},
	}

	for _, test := range tests {
		if result := shiftLastElement(test.input); !equalsSlices(result, test.output) {
			t.Errorf("shiftLastElement(%q) = %v", test.input, test.output)
		}
	}
}

//сравнение двух двумерных слайсов
func equalsSlices(a, b [][]string) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {

		if (a[i][0] != b[i][0]) || (a[i][1] != b[i][1]) {
			return false
		}
	}

	return true
}
