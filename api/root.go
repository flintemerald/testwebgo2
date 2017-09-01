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

  http.Handle("/testapi/test1/", indexTest1())
}



func Start(cfg Config, listener net.Listener) {

  server := &http.Server{}
  server.ReadTimeout = 60 * time.Second
  server.WriteTimeout = 60 * time.Second
  server.MaxHeaderBytes = 1 << 16

  routeUrls(cfg)

  go server.Serve(listener)
}
