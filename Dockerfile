FROM golang:alpine

WORKDIR /app

RUN apk update && apk add gcc g++

run go install github.com/cucumber/godog/cmd/godog@latest

COPY . .

CMD ["go", "run", "cmd/api/main.go"]