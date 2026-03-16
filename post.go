package main

import (
	"fmt"
	"os"
)

type Post struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

func RecordToPost(file string, record Record) (Post, error) {
	data, err := os.ReadFile("posts/" + file + ".md")
	if err != nil {
		panic(err)
	}

	date := fmt.Sprintf("%d-%02d-%02d", record.createdAt.Local().Year(), record.createdAt.Local().Month(), record.createdAt.Local().Day())
	return Post{Title: record.title, Date: date, Content: string(data)}, nil
}
