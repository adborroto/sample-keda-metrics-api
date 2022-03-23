FROM golang:1.17


WORKDIR /app
COPY go.mod .
COPY server.go server.go

RUN go mod tidy
RUN go build -o server server.go

EXPOSE 8090

CMD ["/app/server"]