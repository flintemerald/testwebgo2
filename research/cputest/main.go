package main

import (
  "fmt"
  "math/rand"
  "time"
  "crypto/sha1"
  "encoding/hex"
  "sync"
  "runtime"
)

func RandStringRunes(work_index int, n int, wg *sync.WaitGroup) string {
  defer wg.Done()
  var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
  b := make([]rune, n)
  for i := range b {
      b[i] = letterRunes[rand.Intn(len(letterRunes))]
  }
  randString := string(b)
  sum := sha1.Sum([]byte(randString))
  hexsum := hex.EncodeToString(sum[:])
  fmt.Println(work_index, hexsum)
  return hex.EncodeToString(sum[:])
}

func main() {
  each_work_weight := 9000000
  total_works_count := 10
  cpu_cores_count := 0 // 0 - don't change default behaviour
  
  rand.Seed(time.Now().UnixNano())
  if cpu_cores_count != 0 {
    runtime.GOMAXPROCS(cpu_cores_count)
  }
  wg := sync.WaitGroup{}
  t0 := time.Now()
  for i := 0; i < total_works_count; i+=1 {
    wg.Add(1)
    go RandStringRunes(i, each_work_weight, &wg)
  }
  wg.Wait()
  dt := time.Since(t0)
  fmt.Println("done!", dt)
}
