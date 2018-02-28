package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func routines() {
	start := time.Now()
	for _, url := range webList {
		go getURL(url)
	}
	time.Sleep(10 * time.Second)
	elapsed := time.Since(start)
	log.Println("\n===================================\n Completed in:", elapsed, "\n===================================\n")
}

func getURL(url string) {
	a, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url, " - ", a.Status)
}
