FROM golang:latest

LABEL maintainer="Espoir Ditekemena <espopir@ditkay.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3002

ADD db/migrations/* /app/db/migrations/

CMD ["./main"]