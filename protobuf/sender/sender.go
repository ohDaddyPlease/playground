package main

import (
	"bytes"
	"log"
	"net/http"

	email "github.com/ohdaddyplease/playground/protobuf/emailpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	log.Println("I am have been started")

	e := &email.Email{
		Sender: &email.Person{
			Name:    "Sergey",
			Address: "s@email.com",
		},
		Receiver: &email.Person{
			Name:    "Andrey",
			Address: "a@email.com",
		},
		Message: "Hello, Andrey!",
	}
	log.Printf("email struct: %+v\n", e)

	mmsg, err := proto.Marshal(e)
	if err != nil {
		log.Fatal("marshal: ", err)
	}
	log.Printf("bytes seq (marshaled): %+v\n", mmsg)

	var msg bytes.Buffer
	if _, err := msg.Write(mmsg); err != nil {
		log.Fatal("write to buffer: ", err)
	}
	log.Printf("reader: %+v\n", msg)

	if _, err := http.Post("http://localhost:1111", "application/x-protobuf", &msg); err != nil {
		log.Fatal("request: ", err)
	}
}
