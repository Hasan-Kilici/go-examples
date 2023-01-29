package main

import (
 "database/sql"
 _ "github.com/go-sql-driver/mysql"
)

func main(){
  db, err := sql.Open("mysql", "user:password@tcp(hostname:port)/dbname")
  dbrows, err := trstokdb.Query("SELECT * FROM ogrenci")
  for stokrows.Next() {
    var isim string
    var soyisim string
    var no int
    var sinif string
    
    err = dbrows.Scan(&isim , &soyisim , &no, &sinif)
    if err != nil {
        panic(err.Error())
    }
 }
} 
