package model

import "fmt"

type Topic struct {
	Id int
	Title string
	Content string
}

func GetTopicById(id int) *Topic {
	var topic Topic
	err := getDB().QueryRow("select id,title,content from topic where id=?", id).Scan(&topic.Id, &topic.Title, &topic.Content)
	if err != nil {
		return nil
	}
	return &topic
}

func GetTopics(start, size int) []*Topic {
	var topics []*Topic

	statement, err := getDB().Query("select id, title from topic order by id desc limit ?, ?", start, size)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer statement.Close()
	for statement.Next() {
		var t Topic
		err := statement.Scan(&t.Id, &t.Title)
		if err != nil {
			break
		}
		topics = append(topics, &t)
	}
	return topics
}
