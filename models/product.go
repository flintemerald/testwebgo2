package models

import (
)



type ProductCompact struct {
  Id int64
  Name_rus string
  Name_eng string
  Price int64
}



type ProductFull struct {
  ProductCompact
  
  Photo_url string
  Year string
  Vendor_code string
  Country string
  Alc_vol int32
  Capacity int32
  Description string
  Grape_sort string
}
