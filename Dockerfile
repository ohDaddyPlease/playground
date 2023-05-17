FROM golang:1.20

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -v -o ./app

EXPOSE 80

ENTRYPOINT ["./app"]
