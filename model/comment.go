package model

import "sync"

type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func GetComments(id, lastId int) ([]*Comment, int) {
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
	return comments, total
}

func CreateComment(mapId int, content string) bool {
	statement, err := getDB().Prepare("insert into comments(map_id,content) values(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()
	res, err := statement.Exec(mapId, content)
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