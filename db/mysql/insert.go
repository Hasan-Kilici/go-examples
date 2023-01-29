package main

import (
 "database/sql"
 _ "github.com/go-sql-driver/mysql"
 "fmt"
)

func main() {
  db, err := sql.Open("mysql", "user:password@tcp(hostname:port)/dbname")
  res, err := db.Exec("INSERT INTO Ogrenci (isim, soyisim, sinif, no) VALUES (?, ?, ?, ?)", "Hasan", "KILICI", "12/D", 1200)
  if err != nil {
    fmt.Println(err)
    return
  }

  rowCount, err := res.RowsAffected()
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("%d satır eklendi\n", rowCount)
}
