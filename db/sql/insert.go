package main

import (
	"database/sql"
	"fmt"
)

func main() {
  db, err := sql.Open("sql", "user:password@/dbname");
	res, err := db.Exec("INSERT INTO Ogrenci (isim, soyisim, sinif, no) VALUES (?, ?)", "Hasan", "KILICI", "12/D", 1200)
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

	fmt.Printf("%d satır eklendi\n", rowCount)
}
