.PHONY: build
build: receiver sender

.PHONY: receiver
receiver:
	go build -o ./bin/receiver ./receiver

.PHONY: sender
sender:
	go build -o ./bin/sender ./sender

.PHONY: generate
generate:
	protoc --go_out=./emailpb --go_opt=paths=source_relative email.proto

.PHONY: clean
clean: 
	rm -rf ./bin/receiver ./bin/sender ./emailpb/*