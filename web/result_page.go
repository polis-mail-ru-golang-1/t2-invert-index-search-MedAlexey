package web

import (
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/invertIndex"
	"html/template"
	"net/http"
	"strconv"
)

var InvertIndexMap map[string]map[string]int
var SliceOfNames []string

type Match struct {
	File   string
	Number int
}

func ResultPage(w http.ResponseWriter, r *http.Request) {
	phrase := r.URL.Query().Get("phrase")

	sFullMatches, sNotFullMatches := invertIndex.FindMatches(phrase, InvertIndexMap, SliceOfNames)
	fullMatches, notFullMatches := convertMatchesToStruct(sFullMatches, sNotFullMatches)

	tpl := template.Must(template.ParseFiles("web/result.html"))
	tpl.Execute(w, struct {
		FullIsNotEmpty    bool
		FullMatch         []Match
		NotFullIsNotEmpty bool
		NotFullMatch      []Match
	}{
		!isEmpty(fullMatches),
		fullMatches,
		!isEmpty(notFullMatches),
		notFullMatches,
	})
}

func isEmpty(match []Match) bool {
	if len(match) == 0 {
		return true
	}
	return false
}

// преобразование [][]string в []Match
func convertMatchesToStruct(sFullMatches [][]string, sNotFullMatches [][]string) ([]Match, []Match) {
	var fullMatches []Match
	var notFullMatches []Match

	for _, match := range sFullMatches {
		number, _ := strconv.Atoi(match[1])
		newElement := Match{match[0], number}
		fullMatches = append(fullMatches, newElement)
	}

	for _, match := range sNotFullMatches {
		number, _ := strconv.Atoi(match[1])
		newElement := Match{match[0], number}
		notFullMatches = append(notFullMatches, newElement)
	}

	return fullMatches, notFullMatches
}
