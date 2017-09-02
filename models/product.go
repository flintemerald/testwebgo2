package models

import (
  "github.com/jmoiron/sqlx"
)



type Product struct {
  Id int64
  Name_rus string
  Name_eng string
  Photo_url string
  Year string
  Vendor_code string
  Country string
  Alc_vol int32
  Capacity int32
  Description string
  Grape_sort string
  Price int64
}

func (o *Product) Scan(rows *sqlx.Rows) {
  rows.StructScan(&o)
}
