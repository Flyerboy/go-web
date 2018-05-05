package controllers

import (
	"html/template"
	"net/http"
	"strings"
	"fmt"
	"strconv"
)

func unescaped(x string) interface{} {
	return template.HTML(x)
}

func Render(w http.ResponseWriter, name string, data map[string]interface{}) {

	tplname := strings.TrimRight(name, ".html") + ".html"

	files := []string{
		"header.html",
		tplname,
		"footer.html",
	}

	for k, v := range files{
		files[k] = "template/" + v
	}

	t := template.New("")
	t = t.Funcs(template.FuncMap{"html": unescaped})

	tmpl, err := t.ParseFiles(files[0], files[1], files[2])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func Page(total, size int, r *http.Request) (int, string) {
	p := r.FormValue("p")
	if p == "" {
		p = "1"
	}
	page, _ := strconv.Atoi(p)

	start := (page - 1) * size

	totalPage := total / size

	//show := 5

	html := "<ul class='pagination mt30'>"

	for i := 1; i <= totalPage; i++ {
		if page == i {
			html += fmt.Sprintf("<li class='page-item active'><a href='?p=%d' class='page-link'>%d</a></li>", i, i)
		} else {
			html += fmt.Sprintf("<li class='page-item'><a href='?p=%d' class='page-link'>%d</a></li>", i, i)
		}
	}

	html += "</ul>"
	return start, html
}
