package main

import (
	"flag"
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
	flag.Parse()
	port := ":8080"
	if flag.Arg(0) != "" {
		port = flag.Arg(0)
	}
	http.HandleFunc("/", hello)
	log.Printf("Launching server on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
