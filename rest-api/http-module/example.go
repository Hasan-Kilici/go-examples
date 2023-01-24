package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        users := []User{
            {Name: "John", Age: 30},
            {Name: "Mike", Age: 45},
        }

        json, _ := json.Marshal(users)

        w.Header().Set("Content-Type", "application/json")
        w.Write(json)
    })

    fmt.Println("Server starting...")
    http.ListenAndServe(":8080", nil)
}
