package db

import (
  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"
)



type Config struct {
  ConnectString string
}



var con *sqlx.DB



func Init(cfg Config) {
  var err error
  con, err = sqlx.Open("postgres", cfg.ConnectString)
  if err != nil {
    panic(err)
  }
  err = con.Ping()
  if err != nil {
    panic(err)
  }
}
