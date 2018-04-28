package main

import (
	"net/http"
	"project/controllers"
)

func main() {

	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/", controllers.TopicIndex)
	http.HandleFunc("/topic", controllers.TopicShow)
	http.HandleFunc("/topic/create", controllers.TopicCreate)
	http.HandleFunc("/login", controllers.Login)

	http.ListenAndServe(":8005", nil)

}
