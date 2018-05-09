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
	var comments []*Comment

	ch := make(chan *Comment, 5)
	sy := sync.WaitGroup{}
	sy.Add(2)
	go func(commentChan chan *Comment, sy *sync.WaitGroup) {
		statement, err := getDB().Query("select id,content,created_at from comments where map_id=? order by id desc limit 20", id)
		defer statement.Close()
		if err == nil {

			for statement.Next() {
				var c Comment
				err := statement.Scan(&c.Id, &c.Content, &c.CreatedAt)
				if err != nil {
					break
				}
				commentChan <- &c
			}
		}
		sy.Done()
	}(ch, &sy)

	for i := 0; i < 20; i++ {
		c := <- ch
		comments = append(comments, c)
	}
	close(ch)

	total := 0
	go func(total *int, sy *sync.WaitGroup) {
		err := getDB().QueryRow("select count(id) num from comments where map_id=? limit 1", id).Scan(total)
		if err != nil {
		}
		sy.Done()
	}(&total, &sy)

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