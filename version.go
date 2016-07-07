package main

import (
	"fmt"
	"log"
	"os"
)

type Version struct {
}

func generateInitFolder() {
	dir, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v", dir)
}
