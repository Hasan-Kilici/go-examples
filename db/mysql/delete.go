package main

import (
 "database/sql"
 _ "github.com/go-sql-driver/mysql"
 "fmt"
)

func main() {
  db, err := sql.Open("mysql", "user:password@tcp(hostname:port)/dbname")
  res, err := db.Exec("DELETE FROM Ogrenci WHERE no=1200")
  if err != nil {
    fmt.Println(err)
    return
  }

  // Etkileşimli satır sayısını alın
  rowCount, err := res.RowsAffected()
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("%d satır Silindi\n", rowCount)
}
