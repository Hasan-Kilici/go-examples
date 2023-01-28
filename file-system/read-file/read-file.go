package main

import (
	"fmt"
	"os"
)

func main() {
	read, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(read))
	}
}
