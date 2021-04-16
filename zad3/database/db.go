package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

type Database struct {
	db *sql.DB
}

var instance *Database

var once sync.Once

func GetInstance() *Database {
	once.Do(func() {
		instance = &Database{}
	})
	return instance
}

func (d *Database) Open() {
	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		log.Panic("Cannot access database.")
	}

	d.db = db
}

func (d *Database) Close() {
	err := d.db.Close()

	if err != nil {
		log.Panic("Cannot close database.")
	}
}

func (d *Database) Query(query string) (*sql.Rows, error) {
	rows, err := d.db.Query(query)

	if err != nil {
		log.Println("Query returned error.")
		return nil, err
	}
	return rows, nil
}

func (d *Database) Execute(query string) (int64, error) {
	res, err := d.db.Exec(query)

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func Init() {
	db := GetInstance()
	db.Open()
	db.Execute("CREATE TABLE IF NOT EXISTS UserS (ID integer primary key, NAME TEXT, SURNAME TEXT, BIRTHDAY TEXT)")
	db.Close()
}
