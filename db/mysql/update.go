package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "user:password@tcp(host:port)/dbname")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
	
    res, err := db.Exec("UPDATE Ogrenci SET isim='Hasan', soyisim='KILICI', sinif='12/D' WHERE no=1200")
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

	fmt.Printf("%d satır düzenlendi\n", rowCount)
}
