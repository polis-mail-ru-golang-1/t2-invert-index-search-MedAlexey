package web

import (
	"html/template"
	"net/http"
)

func SearchPage(w http.ResponseWriter, r *http.Request) {

	phrase := r.FormValue("phrase")
	if phrase != "" {
		http.Redirect(w, r, "/result?phrase="+phrase, http.StatusFound)
	}

	tpl := template.Must(template.ParseFiles("web/search.html"))
	tpl.Execute(w, nil)
}
