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



// func peopleHandler(m *model.Model) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		people, err := m.People()
// 		if err != nil {
// 			http.Error(w, "This is an error", http.StatusBadRequest)
// 			return
// 		}
// 
// 		js, err := json.Marshal(people)
// 		if err != nil {
// 			http.Error(w, "This is an error", http.StatusBadRequest)
// 			return
// 		}
// 
// 		fmt.Fprintf(w, string(js))
// 	})
// }
