package main

import (
	"github.com/ruancaetano/go-gin-fruits/internal/app"

	docs "github.com/ruancaetano/go-gin-fruits/docs"
)

func main() {
	docs.SwaggerInfo.Title = "Fruit Crud"
	docs.SwaggerInfo.Version = "1.0"

	s := app.NewServer()
	s.Start()
}
