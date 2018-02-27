package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
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
