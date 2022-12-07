FROM golang:alpine

WORKDIR /app

RUN apk update && apk add gcc g++

COPY . .

CMD ["go", "run", "cmd/api/main.go"]