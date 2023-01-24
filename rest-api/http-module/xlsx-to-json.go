package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/tealeg/xlsx"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./my-app/build")))
    http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }
        file, err := xlsx.OpenFile("urun.xlsx")
        if err != nil {
            panic(err)
        }

        sheet := file.Sheets[0]
        rows := make([]map[string]string, 0)
        for _, row := range sheet.Rows {
            rowMap := make(map[string]string)
            for j, cell := range row.Cells {
                rowMap[sheet.Rows[0].Cells[j].String()] = cell.String()
            }
            rows = append(rows, rowMap)
        }

        data, err := json.Marshal(rows)
        if err != nil {
            panic(err)
        }

        fmt.Println(string(data))
        w.Header().Set("Content-Type","application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(data)
     })
  
    fmt.Println("Server starting...")
    http.ListenAndServe(":8080", nil)
}
