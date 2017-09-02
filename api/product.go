package api

import (
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"
  
  "../utils"
  "../db"
  "../models"
)


func getProductsListApi() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    resDict := make(map[string]interface{})
    resDict["answer"] = db.GetProductsList()
    resDict["status"] = "ok"
    resString, _ := json.Marshal(resDict)
    fmt.Fprintf(w, string(resString))
  })
}



func getProductApi() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    body, _ := ioutil.ReadAll(r.Body)
    dict := make(map[string]interface{})
    json.Unmarshal(body, &dict)
    product_id := utils.GetInt64FromDict(dict, "product_id")
    resDict := make(map[string]interface{})
    resDict["status"] = "ok"
    resDict["answer"] = db.GetProduct(product_id)
    resString, _ := json.Marshal(resDict)
    fmt.Fprintf(w, string(resString))
  })
}



func saveProductApi() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    body, _ := ioutil.ReadAll(r.Body)
    dict := make(map[string]interface{})
    json.Unmarshal(body, &dict)
    dictProd := dict["product"].(map[string]interface{})
    prod := models.ProductFull{}
    prod.Name_rus = dictProd["Name_rus"].(string)
    prod.Name_eng = dictProd["Name_eng"].(string)
    prod.Photo_url = dictProd["Photo_url"].(string)
    prod.Year = dictProd["Year"].(string)
    prod.Vendor_code = dictProd["Vendor_code"].(string)
    prod.Country = dictProd["Country"].(string)
    prod.Alc_vol = utils.GetInt32FromDict(dictProd, "Alc_vol")
    prod.Capacity = utils.GetInt32FromDict(dictProd, "Capacity")
    prod.Description = dictProd["Description"].(string)
    prod.Grape_sort = dictProd["Grape_sort"].(string)
    prod.Price = utils.GetInt64FromDict(dictProd, "Price")
    if _, ok := dict["update_product_id"]; ok {
      update_product_id := utils.GetInt64FromDict(dict, "update_product_id")
      db.UpdateProduct(update_product_id, prod)
    } else {
      db.AddNewProduct(prod)
    }
    resDict := make(map[string]interface{})
    resDict["status"] = "ok"
    resString, _ := json.Marshal(resDict)
    fmt.Fprintf(w, string(resString))
  })
}



func deleteProductApi() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    body, _ := ioutil.ReadAll(r.Body)
    dict := make(map[string]interface{})
    json.Unmarshal(body, &dict)
    product_id := utils.GetInt64FromDict(dict, "product_id")
    db.DeleteProduct(product_id)
    resDict := make(map[string]interface{})
    resDict["status"] = "ok"
    resString, _ := json.Marshal(resDict)
    fmt.Fprintf(w, string(resString))
  })
}
