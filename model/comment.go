package model

import (
	"sync"
	"time"
)

type Comment struct {
	Id int `json:"id"`
	MapId int
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func (this *Comment) GetLists(id, lastId int) ([]*Comment, int) {


	ch := make(chan *Comment, 5)
	sy := sync.WaitGroup{}
	sy.Add(2)
	go func() {
		statement, err := getDB().Query("select id,content,created_at from comments where map_id=? order by id desc limit 20", id)
		defer statement.Close()
		if err == nil {

			for statement.Next() {
				var c Comment
				err := statement.Scan(&c.Id, &c.Content, &c.CreatedAt)
				if err != nil {
					break
				}
				ch <- &c
			}
			close(ch)
		}
		sy.Done()
	}()

	var comments []*Comment
	for i := range ch {
		comments = append(comments, i)
	}

	total := 0
	go func() {
		err := getDB().QueryRow("select count(id) num from comments where map_id=? limit 1", id).Scan(&total)
		if err != nil {
		}
		sy.Done()
	}()

	sy.Wait()
	getDB().Close()
	return comments, total
}

func (this *Comment) Create() bool {
	statement, err := getDB().Prepare("insert into comments(map_id,content,created_at) values(?, ?, ?)")
	defer getDB().Close()
	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()
	res, err := statement.Exec(this.MapId, this.Content, time.Now().Format("2012-01-01 00:00:00"))
	if err != nil {
		return false
	}
	row, err := res.RowsAffected()
	if err != nil {
		return false
	}
	if row > 0 {
		return true
	}
	return false
}