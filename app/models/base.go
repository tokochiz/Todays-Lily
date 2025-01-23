package models

import (

"database/sql"
"log"
"todo_app/config"
"github.com/mattn/go-sqlite3"
)
var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		panic(err)
	}

	cmdU := fmt.Sprif(`CREATE TABLE IF NOT EXISTS %s(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid STRING NOT NULL UNIQUE, 
	name STRING,
	email STRING,
	password STRING,
	created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)
}

func createUUID() (uuid string) {
	uuid = uuid.New().String()
	return
}

func Encrypt(plainText string) (cryptText string) {
	cryptText = fmt.Sprintf("%x", sha1.Sum([]byte(plainText)))
	return
}