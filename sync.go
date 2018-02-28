package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func syncFunc() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	for _, url := range webList {
		wg.Add(1)
		go sync_getURL(url, wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Println("\n===================================\n Completed in:", elapsed, "\n===================================\n")
}

func sync_getURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	a, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println(url, " - ", a.Status)
}
