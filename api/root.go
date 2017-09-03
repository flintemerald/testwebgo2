package api

import (
  // "encoding/json"
  "net"
  "net/http"
  "time"
)



type Config struct {
  StaticPath string
}



func routeUrls(cfg Config) {
  http.Handle("/static3/", http.StripPrefix("/static3/", http.FileServer(http.Dir(cfg.StaticPath))))
  http.Handle("/", indexHandler())

  http.Handle("/testapi3/test1/", test1Api())
  http.Handle("/testapi3/get_products_list/", getProductsListApi())
  http.Handle("/testapi3/get_product/", getProductApi())
  http.Handle("/testapi3/save_product/", saveProductApi())
  http.Handle("/testapi3/delete_product/", deleteProductApi())
}



func Start(cfg Config, listener net.Listener) {

  server := &http.Server{}
  server.ReadTimeout = 60 * time.Second
  server.WriteTimeout = 60 * time.Second
  server.MaxHeaderBytes = 1 << 16

  routeUrls(cfg)

  go server.Serve(listener)
}
