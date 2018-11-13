package web

import (
	"github.com/rs/zerolog"
	"html/template"
	"net/http"
)

var PageLogger zerolog.Logger

var searchTemplate = template.Must(template.ParseFiles("web/layout.html", "web/search.html"))

func SearchPage(w http.ResponseWriter, r *http.Request) {

	printLog(r)

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

func printLog(r *http.Request) {
	PageLogger.Debug().
		Str("method", r.Method).
		Str("remote", r.RemoteAddr).
		Str("path", r.URL.Path).
		Str("phrase", r.URL.Query().Get("phrase")).
		Msgf("Called url %s", r.URL.Path)
}
