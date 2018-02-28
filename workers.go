package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func worker() {
	start := time.Now()

	wg := new(sync.WaitGroup)
	lCH := make(chan string)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go getURLWorker(i, lCH, wg)
	}

	for _, url := range webList {
		lCH <- url
	}
	close(lCH)
	wg.Wait()
	elapsed := time.Since(start)
	log.Println("\n===================================\n Completed in:", elapsed, "\n===================================\n")
}

func getURLWorker(i int, lCH chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range lCH {
		a, err := http.Get(url)
		if err != nil {
			fmt.Println("Worker", i, "ERROR: ", err)
			return
		}
		fmt.Println("Worker", i, url, " - ", a.Status)
	}
}
