package db

import (
  "../models"
)



func GetProductsList() []models.ProductCompact {
  res := []models.ProductCompact{}
  rows, err := con.Queryx("SELECT id, name_rus, name_eng FROM product LIMIT 9000")
  if err != nil {
    return res
  }
  for rows.Next() {
    prod := models.ProductCompact{}
    rows.StructScan(&prod)
    res = append(res, prod)
  }
  return res
}
