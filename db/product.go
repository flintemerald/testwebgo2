package db

import (
  "../models"
)



func GetProductsList() []models.ProductCompact {
  res := []models.ProductCompact{}
  rows, err := con.Queryx("SELECT id, name_rus, name_eng, price FROM product ORDER BY id ASC LIMIT 9000")
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



func GetProduct(product_id int64) models.ProductFull {
  row := con.QueryRowx("SELECT * FROM product WHERE id=$1 LIMIT 1", product_id)
  prod := models.ProductFull{}
  row.StructScan(&prod)
  return prod
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



func UpdateProduct(update_product_id int64, prod models.ProductFull) {
  con.Exec(`
    UPDATE product SET
      photo_url=$1,
      year=$2,
      vendor_code=$3,
      country=$4,
      alc_vol=$5,
      capacity=$6,
      name_rus=$7,
      name_eng=$8,
      description=$9,
      grape_sort=$10,
      price=$11
    WHERE Id=$12`,
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
    prod.Price,
    update_product_id)
}



func DeleteProduct(prodId int64) {
  con.Exec("DELETE FROM product WHERE id=$1", prodId)
}
