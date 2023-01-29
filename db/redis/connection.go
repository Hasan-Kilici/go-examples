package main

import (
 "github.com/go-redis/redis/v8"
)

func main(){
  client := redis.NewClient(&redis.Options{
    Addr:     "hostname:port",
    Password: "password",
    DB:       0,
  })

  val, err := client.Get("key").Result()
  if err != nil {
    panic(err)
  }
  fmt.Println("key", val)
}
