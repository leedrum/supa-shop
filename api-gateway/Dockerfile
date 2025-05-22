FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o gateway main.go

EXPOSE 8080

CMD ["./gateway"]
