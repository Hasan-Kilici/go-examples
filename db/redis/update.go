package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:     "host:port",
        Password: "password",
        DB:       0,
    })
    defer client.Close()
	
    err := client.HMSet("ogrenci:1200", map[string]interface{}{
        "isim":    "Hasan",
        "soyisim": "KILICI",
        "sinif":   "12/D",
    }).Err()
    if err != nil {
        fmt.Println(err)
        return
    }
	
    fmt.Println("Ogrenci bilgileri düzenlendi")
}
