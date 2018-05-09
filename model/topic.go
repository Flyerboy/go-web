package model

import (
	"sync"
	"fmt"
)

type Topic struct {
	Id int
	Title string
	Content string
}

func (this *Topic) GetById(id int) *Topic {
	var topic Topic
	err := getDB().QueryRow("select id,title,content from topic where id=?", id).Scan(&topic.Id, &topic.Title, &topic.Content)
	if err != nil {
		return nil
	}
	return &topic
}

func (this *Topic) GetLists(start, size int) ([]*Topic, int) {
	var topics []*Topic

	ch := make(chan *Topic, 5)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(topicChan chan *Topic, wg *sync.WaitGroup) {
		statement, err := getDB().Query("select id, title from topic order by id desc limit ?, ?", start, size)
		defer statement.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		for statement.Next() {
			var t Topic
			err := statement.Scan(&t.Id, &t.Title)
			if err != nil {
				break
			}
			topicChan <- &t
		}
		wg.Done()
	}(ch, &wg)

	for i := 0; i < size; i++ {
		topics = append(topics, <-ch)
	}
	close(ch)

	total := 0
	go func(total *int, wg *sync.WaitGroup) {
		err := getDB().QueryRow("select count(id) from topic where status=1 limit 1").Scan(total)
		if err != nil {
			// 记录日志
		}
		wg.Done()
	}(&total, &wg)

	wg.Wait()
	return topics, total
}

func (this *Topic) Count() int {
	total := 0
	err := getDB().QueryRow("select count(id) from topic where status=1 limit 1").Scan(&total)
	if err != nil {

	}
	return total
}
