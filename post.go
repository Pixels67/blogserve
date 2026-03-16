package main

import (
	"fmt"
	"strings"
)

type Post struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

func RecordToPost(file string, record Record) (Post, error) {
	date := fmt.Sprintf("%d-%02d-%02d", record.createdAt.Local().Year(), record.createdAt.Local().Month(), record.createdAt.Local().Day())
	slug := strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(record.title), " ", "-"), "#", "")
	return Post{Title: record.title, Slug: slug, Date: date, Content: record.content}, nil
}
