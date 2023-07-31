FROM golang:latest AS builder

WORKDIR /app

COPY go.sum go.mod ./

RUN go mod download 

COPY . .

RUN go build -o app main.go

FROM golang:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 3000

ENV DATABASE_DIR=/data

ENV DATABASE_FILE=database.db

CMD ["./app", "-db", "/data/${DATABASE_FILE}"]
