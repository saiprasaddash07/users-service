FROM golang:latest

WORKDIR /users-service

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

CMD ["./users-service"]