package utils

import (
  "math/rand"
  "strconv"
  "fmt"
)



var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}



func GetInt32FromDict(dict map[string]interface{}, name string) int32 {
  res, _ := strconv.ParseInt(dict[name].(string), 10, 32)
  return int32(res)
}



func GetInt64FromDict(dict map[string]interface{}, name string) int64 {
  unk := dict[name]
  var s string
  switch unk.(type) {
  case string:
    s = unk.(string)
  case float64:
    s = fmt.Sprintf("%.0f", unk.(float64))
  }
  res, _ := strconv.ParseInt(s, 10, 64)
  return int64(res)
}
