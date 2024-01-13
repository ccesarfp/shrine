package main

import (
	"log"
	"net/http"
	"time"
)

const addr string = "8080"

func main() {
	start := time.Now()

	log.Println("Starting application on port", addr)

	srv := &http.Server{
		Addr: ":" + addr,
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Println("Application initialization took", elapsed)

	err := srv.ListenAndServe()
	log.Fatal(err)
}
