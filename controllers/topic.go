package controllers

import (
	"net/http"
	"fmt"
	"project/model"
	"strconv"
	"strings"
)


func TopicShow(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	tid, _ := strconv.Atoi(id)
	data := make(map[string]interface{})
	data["topic"] = model.GetTopicById(tid)
	Render(w, "topic/show", data)
}

func TopicIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := r.Form.Get("p")
	if p == "" {
		p = "1"
	}
	page, _ := strconv.Atoi(p)

	data := make(map[string]interface{})
	size := 10
	data["categories"] = model.GetHotCategory(3)
	start, page_html := Page(100, size, page)
	data["topics"] = model.GetTopics(start, size)
	data["page"] = page_html
	Render(w, "topic/index", data)
}

func TopicCreate(w http.ResponseWriter, r *http.Request) {
	if strings.Compare(r.Method, "POST") == 0 {
		r.ParseForm()
		title := r.PostForm.Get("title")
		content := r.PostForm.Get("content")
		fmt.Println(title, content)
	} else {
		Render(w, "topic/create", nil)
	}
}