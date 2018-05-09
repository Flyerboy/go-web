package controllers

import (
	"net/http"
	"strconv"
	"project/model"
)


func CommentIndex(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	cid, _ := strconv.Atoi(id)
	last := req.FormValue("last_id")
	lastId, _ := strconv.Atoi(last)

	commentModel := model.Comment{}

	comments, total := commentModel.GetLists(cid, lastId)

	res := JsonResponse{
		writer: w,
		Status: http.StatusOK,
		Data: comments,
		Total: total,
		Msg: "ok",
	}
	res.Write()
}

func CommentCreate(w http.ResponseWriter, req *http.Request) {
	mapId := req.FormValue("map_id")
	content := req.FormValue("content")
	id, _ := strconv.Atoi(mapId)

	if id == 0 || content == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentModel := model.Comment{
		MapId: id,
		Content: content,
	}
	err := commentModel.Create()
	if err {
		res := JsonResponse{
			writer: w,
			Status: http.StatusOK,
			Msg: "ok",
		}
		res.Write()
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
