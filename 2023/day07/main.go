package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(contents)
}