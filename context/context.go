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

// Connect init context
func Connect() *Context {
	return &Context{Db: connectDb()}
}

func openDb() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres dbname=dynamics_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
