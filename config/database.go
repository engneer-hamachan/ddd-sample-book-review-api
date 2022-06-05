package config

import (
  _ "github.com/mattn/go-sqlite3"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/jinzhu/gorm"
)

var (
  db *gorm.DB
  err error
)

// DB接続 とりあえずSQLiteで
func Connect() *gorm.DB {

  db, err := gorm.Open("sqlite3", "database.db")
  if err != nil {
    panic("failed to connect database")
  }
  return db


}
