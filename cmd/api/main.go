package main

import "github.com/ruancaetano/go-gin-fruits/internal/app"

func main() {
	s := app.NewServer()
	s.Start()
}
