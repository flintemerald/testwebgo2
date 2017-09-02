package api

import (
  "net/http"
  "encoding/json"
  "fmt"
  
  "../db"
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
  })
}
