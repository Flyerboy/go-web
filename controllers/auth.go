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
	controller := Controller{
		writer: w,
		template: "auth/login",
		data: make(map[string]interface{}),
	}

	if strings.Compare(r.Method, "POST") == 0 {
		email := r.FormValue("email")
		password := r.FormValue("password")
		userModel := model.User{
			Email: email,
			Password: password,
		}
		controller.Assign("user", userModel)
		user, err := userModel.Login()
		if err == nil {
			model.SetLogin(w, user)
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}
		controller.Assign("Error", err.Error())
	}

	controller.Render()
}

func Register(w http.ResponseWriter, req *http.Request)  {
	controller := Controller{
		writer: w,
		template: "auth/register",
		data: make(map[string]interface{}),
	}

	if strings.Compare(req.Method, "POST") == 0 {
		name := req.FormValue("name")
		email := req.FormValue("email")
		password := req.FormValue("password")
		rePassword := req.FormValue("repassword")
		userModel := model.User{
			Name: name,
			Email: email,
			Password: password,
		}
		controller.Assign("user", userModel)
		if strings.Compare(password, rePassword) != 0 {
			// 使用 flash
			controller.Assign("Error", "两次密码输入不一致")
			controller.Render()
			return
		} else {
			if userModel.CheckEmail(email) {
				controller.Assign("Error", "邮箱已注册")
				controller.Render()
				return
			}

			res := userModel.Create()
			if res {
				http.Redirect(w, req, "/", http.StatusMovedPermanently)
			}
			controller.Assign("Error", "注册失败")
		}
	} else {
		controller.Assign("Error", "")
	}
	controller.Render()

}
