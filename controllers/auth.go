package controllers

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	Render(w, "auth/login", nil)
}
