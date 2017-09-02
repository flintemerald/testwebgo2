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
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(cfg.StaticPath))))
  http.Handle("/", indexHandler())

  http.Handle("/testapi/test1/", test1Api())
  http.Handle("/testapi/get_products_list/", getProductsListApi())
  http.Handle("/testapi/get_product/", getProductApi())
  http.Handle("/testapi/save_product/", saveProductApi())
  http.Handle("/testapi/delete_product/", deleteProductApi())
}



func Start(cfg Config, listener net.Listener) {

  server := &http.Server{}
  server.ReadTimeout = 60 * time.Second
  server.WriteTimeout = 60 * time.Second
  server.MaxHeaderBytes = 1 << 16

  routeUrls(cfg)

  go server.Serve(listener)
}
