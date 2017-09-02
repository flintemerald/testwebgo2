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



func AddNewProduct(prod models.ProductFull) {
  con.Exec(`
    INSERT INTO product(
      photo_url,
      year,
      vendor_code,
      country,
      alc_vol,
      capacity,
      name_rus,
      name_eng,
      description,
      grape_sort,
      price
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
    prod.Photo_url,
    prod.Year,
    prod.Vendor_code,
    prod.Country,
    prod.Alc_vol,
    prod.Capacity,
    prod.Name_rus,
    prod.Name_eng,
    prod.Description,
    prod.Grape_sort,
    prod.Price)
}



func DeleteProduct(prodId int64) {
  con.Exec("DELETE FROM product WHERE id=$1", prodId)
}
