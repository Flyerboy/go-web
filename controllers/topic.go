package controllers

import (
	"net/http"
	"project/model"
	"strconv"
	"strings"
)


func TopicShow(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	tid, _ := strconv.Atoi(id)
	data := make(map[string]interface{})
	data["topic"] = model.GetTopicById(tid)
	Render(w, "topic/show", data)
}

func TopicIndex(w http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	size := 10
	data["categories"] = model.GetHotCategory(3)
	start, page_html := Page(100, size, req)
	data["topics"] = model.GetTopics(start, size)
	data["page"] = page_html
	Render(w, "topic/index", data)
}

func TopicCreate(w http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	if strings.Compare(req.Method, "POST") == 0 {
		title := req.FormValue("title")
		content := req.FormValue("content")

		topic := model.Topic{
			Title: title,
			Content: content,
		}
		data["Topic"] = topic
	}
	Render(w, "topic/create", data)
}