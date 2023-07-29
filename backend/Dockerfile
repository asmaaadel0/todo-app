FROM golang:latest

WORKDIR /app

COPY . .


RUN go build -o app cmd/main.go

EXPOSE 3000
CMD [ "./app","-db","database.db" ]