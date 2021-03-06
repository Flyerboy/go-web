package controllers

import (
	"net/http"
	"project/model"
	"strconv"
	"strings"
	"project/utils"
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

	sort := req.FormValue("sort")
	if sort == "" {
		sort = "new"
	}
	controller.data["sorts"] = sort

	size := 10
	topicModel := model.Topic{}
	count := topicModel.Count()
	p := getPage(req)
	start := (p - 1) * size

	page := utils.Page{
		Total: count,
		CurrentPage: p,
		PageSize: size,
	}
	pageHtml := page.ToHTML()
	topics, _ := topicModel.GetLists(start, size)
	controller.Assign("topics", topics)
	controller.Assign("page", pageHtml)

	controller.Assign("user", model.CheckLogin(req))

	categoryModel := model.Category{}
	categories := categoryModel.GetHot(3)
	controller.Assign("categories", categories)
	controller.Render()
}

func TopicCreate(w http.ResponseWriter, req *http.Request) {
	controller := Controller{
		writer: w,
		template: "topic/create",
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