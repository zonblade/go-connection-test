FROM golang:1.20-bookworm

RUN go mod download
RUN go build -ldflags="-s" cmd/main.go -o "db-tester"

CMD ["./db-tester"]