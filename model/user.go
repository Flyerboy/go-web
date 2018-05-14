package model

import (
	"net/http"
	"strconv"
	"fmt"
	"errors"
	"strings"
	"crypto/sha256"
	"io"
)

type User struct {
	Id int
	Name string
	Email string
	Password string
}

const PASSWORD_SALT = "#25%C7"
// 密码加密
func encryptPassword(password string) string {
	pwd := password + PASSWORD_SALT
	h := sha256.New()
	io.WriteString(h, pwd)
	return fmt.Sprintf("%x", h.Sum(nil))
}


// 检查是否登录
func CheckLogin(r *http.Request) *User {
	userId, err := r.Cookie("user_id")
	if err != nil {
		return nil
	}
	id, _ := strconv.Atoi(userId.Value)
	user, _ := GetUser(id)
	return user
}

// 设置登录cookie
func SetLogin(w http.ResponseWriter, user *User)  {
	cookie := http.Cookie{
		Name: "user_id",
		Value: strconv.Itoa(user.Id),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}


// 根据ID获取用户信息
func GetUser(id int) (*User, error) {
	var user User
	err := getDB().QueryRow("select id,name,email from users where id=?", id).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		return &user, err
	}
	return &user, err

}

// 检查邮箱是否存在
func (this *User) CheckEmail(email string) bool {
	count := 0
	err := getDB().QueryRow("select count(id) count from users where email=? limit 1", email).Scan(&count)
	if err == nil && count > 0 {
		return true
	}
	return false
}

// 用户登录
func (this *User) Login() (*User, error)  {
	var user User
	err := getDB().QueryRow("select id,name,email,password from users where email=?", this.Email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil || user.Id == 0 {
		return &user, errors.New("用户不存在")
	}

	password := encryptPassword(this.Password)
	if strings.Compare(user.Password, password) != 0 {
		return &user, errors.New("用户名或密码错误")
	}

	return &user, nil
}

// 创建用户
func (this *User) Create() bool {
	statement, err := getDB().Prepare("insert into users(name,email,password) values(?, ?, ?)")
	defer statement.Close()
	if err != nil {
		panic(err.Error())
	}
	password := encryptPassword(this.Password)
	res, err := statement.Exec(this.Name, this.Email, password)
	if err != nil {
		panic(err.Error())
	}
	row, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	if row > 0 && id > 0 {
		return true
	}
	return false
}


func UpdateUser(id int, user *User) bool {
	statement, err := getDB().Prepare("update users set name=?, email=? where id=?")
	defer statement.Close()
	if err != nil {
		panic(err.Error())
	}
	res, err := statement.Exec(user.Name, user.Email, id)
	if err != nil {
		panic(err.Error())
	}
	row, err := res.RowsAffected()
	if row > 0 {
		return true
	}
	return false
}

// 删除用户
func (this *User) Delete(id int) bool {
	statement, err := getDB().Prepare("delete from users where id=?")
	defer statement.Close()
	if err != nil {
		panic(err.Error())
	}
	res, err := statement.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	row, err := res.RowsAffected()
	if row > 0 {
		return true
	}
	return false
}