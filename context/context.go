package context

import (
	"database/sql"
	"log"
)

// Context project
type Context struct {
	Db *sql.DB
}

// New init context
func New() *Context {
	return &Context{Db: openDb()}
}

func openDb() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres dbname=parking_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
