package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer log.Printf("%v | %v | %s", r.RequestURI, r.UserAgent(), time.Since(start))

	queryValues := r.URL.Query()
	name := queryValues.Get("name")
	if name == "" {
		fmt.Fprintf(w, "Hello Anonymous beast! :)")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func server() {
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
