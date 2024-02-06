package main

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"log"
)

func main() {
	v7, err := uuid.NewV7()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your UUID is:")
	fmt.Println(v7)
}
