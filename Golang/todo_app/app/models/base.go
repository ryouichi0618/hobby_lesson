package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME,
		updated_at DATETIME
	)
	`, tableNameUser)
	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}

	cmdT := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME,
		updated_at DATETIME
	)
	`, tableNameTodo)
	_, err = Db.Exec(cmdT)
	if err != nil {
		log.Fatalln(err)
	}

	cmdS := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME,
		updated_at DATETIME
	)
	`, tableNameSession)
	_, err = Db.Exec(cmdS)
	if err != nil {
		log.Fatalln(err)
	}
}

func createUUID() (uuidObj uuid.UUID) {
	uuidObj, _ = uuid.NewUUID()
	return
}

func Encrypt(plainText string) (crepeText string) {
	crepeText = fmt.Sprintf("%x", sha1.Sum([]byte(plainText)))
	return
}
