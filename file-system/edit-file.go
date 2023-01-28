package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	f.Write([]byte("Merhaba Dünya\n"))
	f.Close()
}
