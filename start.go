package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func simple() {
	start := time.Now()
	for _, url := range webList {
		a, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(url, " - ", a.Status)
	}
	elapsed := time.Since(start)
	log.Println("\n===================================\n Completed in:", elapsed, "\n===================================\n")
}
