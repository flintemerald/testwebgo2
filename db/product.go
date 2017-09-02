package db

import (
  "../models"
)

func GetProductsList() []models.Product {
  res := []models.Product{}
  rows, err := con.Queryx("SELECT id, name_rus, name_eng FROM product LIMIT 9000")
  if err != nil {
    return res
  }
  for rows.Next() {
    prod := models.Product{}
    prod.Scan(rows)
    res = append(res, prod)
  }
  return res
}
