package main

import (
	"net/http"
	"project/controllers"
	"strings"
	"fmt"
)

func main() {

	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/", controllers.TopicIndex)
	http.HandleFunc("/topic/index", controllers.TopicIndex)
	http.HandleFunc("/topic", controllers.TopicShow)
	http.HandleFunc("/topic/create", controllers.TopicCreate)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)
	http.HandleFunc("/register", controllers.Register)

	http.HandleFunc("/comments", controllers.CommentIndex)
	http.HandleFunc("/comment/create", controllers.CommentCreate)

	http.ListenAndServe(":8005", nil)

}

func route(url string) http.Handler {
	arr := strings.Split(url, "@")

	fmt.Println(arr)
	return nil
}
