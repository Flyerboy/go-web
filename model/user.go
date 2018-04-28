package model

import (
	"database/sql"
	"net/http"
	"strconv"
	"fmt"
)

type User struct {
	Id int
	Name string
	Email string
	Password string
}

func GetUserByEmail(email string) *User {
	var user User
	err := DB.QueryRow("select id,name,email,password from users where email=?", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return &user
	}
	return &user
}

func CheckLogin(r *http.Request) *User {
	user_id, err := r.Cookie("user_id")
	fmt.Println(user_id)
	if err != nil {
		return nil
	}
	id, _ := strconv.Atoi(user_id.Value)
	user, _ := GetUser(id)
	return user
}

func SetLogin(w http.ResponseWriter, user *User)  {
	cookie := http.Cookie{
		Name: "user_id",
		Value: strconv.Itoa(user.Id),
	}
	http.SetCookie(w, &cookie)
}


func GetUser(id int) (*User, error) {
	/*var user User
	statment, err := DB.Prepare("select id,name,email from users where id=?")
	if err != nil {
		panic(err.Error())
	}
	defer statment.Close()
	row, err := statment.Query(id)
	if row.Next() {
		row.Scan(&user.Id, &user.Name, &user.Email)
	}
	return &user*/

	var user User
	err := DB.QueryRow("select id,name,email from users where id=?", id).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		panic(err.Error())
	}
	return &user, err

}

func SelectUser(id int) []*User {
	db, err := sql.Open("mysql", "root:root@/video")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("select id,name,email from users limit 10")
	if err != nil {
		panic(err.Error())
	}
	/*
	users, err := rows.Columns()
	fmt.Println(users)
	*/
	var users []*User
	//users := make([]*User, 10)
	i := 0

	defer rows.Close()
	for rows.Next() {
		var tmp User
		err := rows.Scan(&tmp.Id, &tmp.Name, &tmp.Email)
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(user)
		//users[i] = &tmp
		users = append(users, &tmp)
		i++
	}
	//fmt.Println(users)
	return users
}

func CreateUser() bool {
	statement, err := DB.Prepare("insert into users(name,email) values(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()
	res, err := statement.Exec("wang", "wang@aa.com")
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
	statement, err := DB.Prepare("update users set name=?, email=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()
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

func DeleteUser(id int) bool {
	statement, err := DB.Prepare("delete from users where id=?")
	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()
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