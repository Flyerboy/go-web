package controllers

import (
	"html/template"
	"net/http"
	"strings"
	"fmt"
	"strconv"
	"encoding/json"
)

type Controller struct {
	writer http.ResponseWriter
	template string
	data map[string]interface{}
}

func unescaped(x string) interface{} {
	return template.HTML(x)
}

func (this *Controller) Assign(name string, value interface{}) {
	this.data[name] = value
}

func (this *Controller) Render() {

	tplname := strings.TrimRight(this.template, ".html") + ".html"

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
		http.Error(this.writer, err.Error(), http.StatusInternalServerError)
	}
	err = tmpl.ExecuteTemplate(this.writer, this.template, this.data)
	if err != nil {
		http.Error(this.writer, err.Error(), http.StatusInternalServerError)
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

type JsonResponse struct {
	writer http.ResponseWriter
	Status int `json:"status"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
	Total int `json:"total"`
}

func (this *JsonResponse)Write() {
	str, err := json.Marshal(this)
	if err == nil {
		this.writer.Header().Set("Content-Type", "application/json")
		this.writer.Write(str)
	} else {
		this.writer.WriteHeader(http.StatusNotFound)
	}

}

