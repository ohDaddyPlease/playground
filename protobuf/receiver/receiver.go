package main

import (
	"io"
	"log"
	"net/http"

	email "github.com/ohdaddyplease/playground/protobuf/emailpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	log.Println("I am have been started")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		e := &email.Email{}
		msg, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal("read msg: ", err)
		}
		r.Body.Close()

		if err := proto.Unmarshal(msg, e); err != nil {
			log.Fatal("unmarshal: ", err)
		}

		log.Printf("bytes seq: %+v", msg)
		log.Printf("unmarshaled msg: %+v", e)
	})
	log.Fatal(http.ListenAndServe(":1111", nil))
}
