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
  
  res, err := client.Del("key").Result()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("%d key silindi\n", res)
}
