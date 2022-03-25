FROM golang:1.17


WORKDIR /app
COPY go.mod .
COPY server.go server.go
COPY auth.go auth.go

RUN go mod tidy
RUN go build -o server .

EXPOSE 8090

CMD ["/app/server"]