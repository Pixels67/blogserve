package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

type Record struct {
	createdAt time.Time
	viewCount int64
	title     string
	content   string
}

func (db *Database) Init() error {
	passwd := os.Getenv("BLOGSERVE_PASSWD")
	connUri := fmt.Sprintf("postgresql://postgres.lghguqmfyfcylpzdjdbo:%s@aws-1-eu-west-1.pooler.supabase.com:5432/postgres", passwd)
	pool, err := pgxpool.New(context.Background(), connUri)
	if err != nil {
		return err
	}

	db.pool = pool
	return nil
}

func (db *Database) Deinit() {
	db.pool.Close()
}

func (db *Database) Get(file string) (Record, error) {
	var record Record

	sql := "select created_at, view_count, title, content from \"Posts\" where file = $1"
	err := db.pool.QueryRow(context.Background(), sql, file).Scan(&record.createdAt, &record.viewCount, &record.title, &record.content)
	if err == nil {
		record.viewCount++
		db.Set(file, record)
	}

	return record, err
}

func (db *Database) GetAll() (map[string]Record, error) {
	records := map[string]Record{}

	sql := "select file, created_at, view_count, title, content from \"Posts\""
	rows, err := db.pool.Query(context.Background(), sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var record Record
		var file string
		rows.Scan(&file, &record.createdAt, &record.viewCount, &record.title, &record.content)
		records[file] = record
	}

	return records, nil
}

func (db *Database) Set(file string, record Record) error {
	sql := "update \"Posts\" set created_at = $1, view_count = $2, title = $3, content = $4 where file = $5"
	_, err := db.pool.Exec(context.Background(), sql, record.createdAt, record.viewCount, record.title, record.content, file)

	return err
}
