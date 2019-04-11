package db

import (
	"database/sql"
	"github.com/fgehrlicher/reminder-bot/conf"
	_ "github.com/mattn/go-sqlite3"
)

func GetConnection(config *conf.Config) (*sql.DB, error) {
	return sql.Open("sqlite3", config.Database.Path)
}
