package main

import (
  "flag"
  "fmt"
  "log"
  "net"
  "os"
  "os/signal"
  "syscall"
  "runtime"
  "math/rand"
  "time"

  "./api"
)



type config struct {
  ListenSpec string
  GoMaxProcs int

  // Db db.Config
  API api.Config
}



func main() {
  fmt.Println("call main.main()...")

  rand.Seed(time.Now().UnixNano())

  cfg := processFlags()
  if cfg.GoMaxProcs != 0 {
    runtime.GOMAXPROCS(cfg.GoMaxProcs)
  }

  if err := run(cfg); err != nil {
    log.Printf("Error in main(): %v", err)
  }

  fmt.Println("call main.main()... done!")
}



func processFlags() *config {
  cfg := &config{}
  flag.StringVar(&cfg.ListenSpec, "listen", "localhost:8080", "HTTP listen spec")
  flag.IntVar(&cfg.GoMaxProcs, "gomaxprocs", 0, "GOMAXPROCS, 0==defalut")
  //flag.StringVar(&cfg.Db.ConnectString, "db-connect", "host=/var/run/postgresql dbname=gowebapp sslmode=disable", "DB Connect String")
  flag.StringVar(&cfg.API.StaticPath, "static-path", "static", "Path to static files dir")

  flag.Parse()
  return cfg
}



func run(cfg *config) error {
  log.Printf("Starting, HTTP on: %s\n", cfg.ListenSpec)

  l, err := net.Listen("tcp", cfg.ListenSpec)
  if err != nil {
    log.Printf("Error creating listener: %v\n", err)
    return err
  }

  api.Start(cfg.API, l)

  waitForSignal()

  return nil
}



func waitForSignal() {
  ch := make(chan os.Signal)
  signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
  s := <-ch
  log.Printf("Got signal: %v, exiting.", s)
}
