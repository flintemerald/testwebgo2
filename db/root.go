package db

import (
  "database/sql"
  _ "github.com/lib/pq"
)



type Config struct {
  ConnectString string
}



var con *sql.DB



func Init(cfg Config) {
  var err error
  con, err = sql.Open("postgres", cfg.ConnectString)
  if err != nil {
    panic(err)
  }
  err = con.Ping()
  if err != nil {
    panic(err)
  }
}
