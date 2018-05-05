package controllers

import (
	"net/http"
	"project/model"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	login := model.CheckLogin(r)
	if login != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	if strings.Compare(r.Method, "POST") == 0 {
		email := r.FormValue("email")
		password := r.FormValue("password")
		user := model.GetUserByEmail(email)
		/*if user.Id > 0 && strings.Compare(user.Password, password) == 0 {
			w.Write([]byte("login success"))
		} else {
			w.Write([]byte("login error"))
		}*/

		if user.Id > 0 && password != "" {
			model.SetLogin(w, user)
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		} else {
			http.Error(w, "登录失败", http.StatusUnauthorized)
		}

	} else {
		Render(w, "auth/login", nil)
	}
}

func Register(w http.ResponseWriter, r *http.Request)  {
	if strings.Compare(r.Method, "POST") == 0 {

	} else {
		Render(w, "auth/register", nil)
	}
}
