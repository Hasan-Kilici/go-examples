package main

import (
	"database/sql"
	"fmt"
)

func main() {
  db, err := sql.Open("sql", "user:password@/dbname");
	res, err := db.Exec("DELETE FROM Ogrenci WHERE no=1200");
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
