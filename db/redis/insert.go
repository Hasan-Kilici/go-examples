package main

import (
 "github.com/go-redis/redis/v8"
 "fmt"
)

func main() {
  client := redis.NewClient(&redis.Options{
    Addr:     "hostname:port",
    Password: "password",
    DB:       0,
  })
  
  res, err := client.HMSet("ogrenci:1200", map[string]interface{}{
    "isim": "Hasan",
    "soyisim": "KILICI",
    "sinif": "12/D",
    "no": "1200",
  }).Result()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("%v eklendi\n", res)
}
