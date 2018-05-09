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

	controller := Controller{
		writer: w,
		template: "topic/show",
		data: make(map[string]interface{}),
	}
	topicModel := model.Topic{}
	topic := topicModel.GetById(tid)
	controller.Assign("topic", *topic)
	controller.Render()
}

func TopicIndex(w http.ResponseWriter, req *http.Request) {

	controller := Controller{
		writer: w,
		template: "topic/index",
		data: make(map[string]interface{}),
	}
	size := 10
	topicModel := model.Topic{}
	count := topicModel.Count()
	start, pageHtml := Page(count, size, req)
	topics, _ := topicModel.GetLists(start, size)
	controller.Assign("topics", topics)
	controller.Assign("page", pageHtml)
	categoryModel := model.Category{}
	categories := categoryModel.GetHot(3)
	controller.Assign("categories", categories)
	controller.Render()
}

func TopicCreate(w http.ResponseWriter, req *http.Request) {
	controller := Controller{
		writer: w,
		template: "topic/index",
		data: make(map[string]interface{}),
	}
	if strings.Compare(req.Method, "POST") == 0 {
		title := req.FormValue("title")
		content := req.FormValue("content")

		topic := model.Topic{
			Title: title,
			Content: content,
		}
		controller.Assign("topic", topic)
	}
	controller.Render()
}