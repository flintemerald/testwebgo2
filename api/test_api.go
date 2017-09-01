package api

import (
  "fmt"
  "net/http"
  _ "time"
  "io/ioutil"
  "encoding/json"
  "strconv"
  "crypto/sha1"
  "encoding/hex"

  "../utils"
)



func indexHandler() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "okoko index text")
  })
}



func indexTest1() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    resDict := make(map[string]interface{})
    body, _ := ioutil.ReadAll(r.Body)
    dict := make(map[string]string)
    json.Unmarshal(body, &dict)
    weight, _ := strconv.ParseInt(dict["weight"], 10, 32)
    //time.Sleep(500 * time.Millisecond)
    randString := utils.RandStringRunes(int(weight))
    sum := sha1.Sum([]byte(randString))
    resDict["answer"] = hex.EncodeToString(sum[:])
    resString, _ := json.Marshal(resDict)
    fmt.Fprintf(w, string(resString))
  })
}