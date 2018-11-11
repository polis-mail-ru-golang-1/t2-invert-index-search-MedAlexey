package web

import (
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-MedAlexey/logger"
	"html/template"
	"net/http"
)

var searchTemplate = template.Must(template.ParseFiles("web/layout.html", "web/search.html"))

func SearchPage(w http.ResponseWriter, r *http.Request) {

	logger.PrintLog("LOG [" + r.Method + "]" + " " + r.RemoteAddr + " " + r.URL.Path + " " + r.URL.Query().Get("phrase"))

	phrase := r.FormValue("phrase")
	if phrase != "" {
		http.Redirect(w, r, "/result?phrase="+phrase, http.StatusFound)
	}

	searchTemplate.ExecuteTemplate(w, "layout", struct {
		Title string
	}{
		"Search",
	})
}
