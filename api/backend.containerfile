FROM golang:1.16

RUN mkdir /app
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
RUN go mod download
COPY server.go /app

RUN go build server.go

CMD ["./server"]