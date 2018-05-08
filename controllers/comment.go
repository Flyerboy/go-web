package controllers

import (
	"net/http"
	"strconv"
	"project/model"
	"encoding/json"
)



func CommentIndex(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	cid, _ := strconv.Atoi(id)
	last := req.FormValue("last_id")
	lastId, _ := strconv.Atoi(last)
	comments, total := model.GetComments(cid, lastId)

	res := JsonResponse{
		Status: 200,
		Msg: "ok",
		Data: comments,
		Total: total,
	}
	str, err := json.Marshal(res)
	if err == nil {
		w.Write(str)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func CommentCreate(w http.ResponseWriter, req *http.Request) {
	mapId := req.FormValue("map_id")
	content := req.FormValue("content")
	id, _ := strconv.Atoi(mapId)

	err := model.CreateComment(id, content)
	if err {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
