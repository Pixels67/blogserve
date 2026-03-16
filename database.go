package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	conn *pgx.Conn
}

type Record struct {
	createdAt time.Time
	viewCount int64
	title     string
}

func (db *Database) Init() error {
	passwd := os.Getenv("BLOGSERVE_PASSWD")
	connUri := fmt.Sprintf("postgresql://postgres:%s@db.lghguqmfyfcylpzdjdbo.supabase.co:5432/postgres", passwd)
	result, err := pgx.Connect(context.Background(), connUri)
	if err != nil {
		return err
	}

	db.conn = result
	return nil
}

func (db *Database) Deinit() error {
	return db.conn.Close(context.Background())
}

func (db *Database) Get(file string) (Record, error) {
	var record Record

	sql := "select created_at, view_count, title from \"Posts\" where file = $1"
	err := db.conn.QueryRow(context.Background(), sql, file).Scan(&record.createdAt, &record.viewCount, &record.title)
	if err == nil {
		record.viewCount++
		db.Set(file, record)
	}

	return record, err
}

func (db *Database) Set(file string, record Record) error {
	sql := "update \"Posts\" set created_at = $1, view_count = $2, title = $3 where file = $3"
	_, err := db.conn.Exec(context.Background(), sql, record.createdAt, record.viewCount, record.title, file)

	return err
}
