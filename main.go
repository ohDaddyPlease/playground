package main

import (
	"net/http"
	"log"
)

func main() {
	log.Println("I am have been started")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("I am alive"))
		if err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
